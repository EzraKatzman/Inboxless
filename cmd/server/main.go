package main

import (
	"fmt"
	"net/http"

	"github.com/EzraKatzman/Inboxless/backend/internal/handlers"
	"github.com/EzraKatzman/Inboxless/backend/internal/redis"
)

func main() {
	redis.InitRedis()

	if pong, err := redis.Rdb.Ping(redis.Ctx).Result(); err != nil {
		panic("Redis connection failed: " + err.Error())
	} else {
		fmt.Println("Redis connected:", pong)
	}

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "pong")
	})

	http.HandleFunc("/api/inbox", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handlers.CreateInboxHandler(w, r)
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/api/inbox/ws", handlers.InboxWebSocketHandler)

	http.HandleFunc("/api/inbox/publish", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		inboxID := r.URL.Query().Get("id")
		message := r.URL.Query().Get("msg")

		if inboxID == "" || message == "" {
			http.Error(w, "Missing ID or message parameters", http.StatusBadRequest)
			return
		}

		channel := fmt.Sprintf("inbox:%s", inboxID)

		err := redis.Rdb.Publish(redis.Ctx, channel, message).Err()
		if err != nil {
			http.Error(w, "Failed to publish message", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Message published"))

		messagesKey := fmt.Sprintf("inbox:%s:messages", inboxID)

		err = redis.Rdb.RPush(redis.Ctx, messagesKey, message).Err()
		if err != nil {
			http.Error(w, "Failed to save message", http.StatusInternalServerError)
			return
		}

		err = redis.Rdb.Expire(redis.Ctx, messagesKey, handlers.InboxTTL).Err()
		if err != nil {
			fmt.Println("Failed to set TTL on messages key:", err)
		}
	})

	http.HandleFunc("/api/inbox/messages", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handlers.GetMessagesHandler(w, r)
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/api/inbox/ttl", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handlers.GetInboxTTLHandler(w, r)
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
