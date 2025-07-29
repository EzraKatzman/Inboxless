package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/EzraKatzman/Inboxless/backend/internal/email"
	"github.com/EzraKatzman/Inboxless/backend/internal/redis"
)

const InboxTTL = 15 * time.Minute

type InboxResponse struct {
	InboxID   string `json:"inbox_id"`
	Email     string `json:"email"`
	ExpiresIn int    `json:"expires_in"` // seconds
}

type TTLResponse struct {
	ExpiresIn int `json:"expires_in"` //seconds
}
type TimedMessage struct {
	Message   string `json:"message"`
	CreatedAt int64  `json:"createdAt"`
}

func CreateInboxHandler(w http.ResponseWriter, r *http.Request) {
	inboxID := email.GenerateInboxId()
	key := fmt.Sprintf("inbox:%s", inboxID)

	err := redis.Rdb.Set(redis.Ctx, key, "[]", InboxTTL).Err()
	if err != nil {
		http.Error(w, "Failed to create inbox", http.StatusInternalServerError)
		return
	}

	resp := InboxResponse{
		InboxID:   inboxID,
		Email:     fmt.Sprintf("%s@inboxless.io", inboxID),
		ExpiresIn: int(InboxTTL.Seconds()),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func GetMessagesHandler(w http.ResponseWriter, r *http.Request) {
	inboxID := r.URL.Query().Get("id")
	if inboxID == "" {
		http.Error(w, "Missing inbox ID", http.StatusBadRequest)
		return
	}

	messagesKey := fmt.Sprintf("inbox:%s:messages", inboxID)
	messages, err := redis.Rdb.LRange(redis.Ctx, messagesKey, 0, -1).Result()
	if err != nil {
		http.Error(w, "Failed to get messages", http.StatusInternalServerError)
		return
	}

	var parsedMessages []TimedMessage
	for _, message := range messages {
		var msg TimedMessage
		if err := json.Unmarshal([]byte(message), &msg); err != nil {
			continue
		}
		parsedMessages = append(parsedMessages, msg)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(parsedMessages)
}

func GetInboxTTLHandler(w http.ResponseWriter, r *http.Request) {
	inboxID := r.URL.Query().Get("id")
	if inboxID == "" {
		http.Error(w, "Missing inbox ID", http.StatusBadRequest)
		return
	}

	ttl, err := GetInboxTTL(inboxID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	resp := TTLResponse{
		ExpiresIn: int(ttl.Seconds()),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func GetInboxTTL(inboxID string) (time.Duration, error) {
	key := fmt.Sprintf("inbox:%s", inboxID)
	ttl, err := redis.Rdb.TTL(redis.Ctx, key).Result()
	if err != nil {
		return 0, err
	}
	if ttl <= 0 {
		return 0, fmt.Errorf("inbox expired or not found")
	}
	return ttl, nil
}
