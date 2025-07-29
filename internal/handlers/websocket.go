package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/EzraKatzman/Inboxless/backend/internal/redis"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// current accepts from all origins
		return true
	},
}

func InboxWebSocketHandler(w http.ResponseWriter, r *http.Request) {
	inboxID := r.URL.Query().Get("id")
	if inboxID == "" {
		http.Error(w, "Missing inbox ID", http.StatusBadRequest)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "WebSocket upgrade failed", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	channel := fmt.Sprintf("inbox:%s", inboxID)
	pubsub := redis.Rdb.Subscribe(ctx, channel)
	defer pubsub.Close()

	for {
		msg, err := pubsub.ReceiveMessage(ctx)
		if err != nil {
			fmt.Println("ReceiveMessage error:", err)
			break
		}

		err = conn.WriteMessage(websocket.TextMessage, []byte(msg.Payload))
		if err != nil {
			fmt.Println("WebSocket write error:", err)
			break
		}

		err = redis.Rdb.Expire(ctx, channel, InboxTTL).Err()
		if err != nil {
			fmt.Println("Failed to refresh inbox TTL on websocket:", err)
		}
	}
}
