package smtp

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/mail"
	"strings"
	"time"

	"github.com/EzraKatzman/Inboxless/backend/internal/handlers"
	"github.com/EzraKatzman/Inboxless/backend/internal/redis"
	"github.com/emersion/go-smtp"
)

type Backend struct{}

type Session struct {
	from string
	to   []string
	data strings.Builder
}

func (b *Backend) Login(state *smtp.Conn, username, password string) (smtp.Session, error) {
	return nil, smtp.ErrAuthUnsupported
}

func (b *Backend) AnonymousLogin(state *smtp.Conn) (smtp.Session, error) {
	log.Println("New anonymous login attempt")
	return &Session{}, nil
}

func (s *Session) Mail(from string, opts *smtp.MailOptions) error {
	log.Printf("MAIL FROM: %s", from)
	s.from = from
	return nil
}

func (s *Session) Rcpt(to string, opts *smtp.RcptOptions) error {
	log.Printf("RCPT TO: %s", to)
	s.to = append(s.to, to)
	return nil
}

func (s *Session) Data(r io.Reader) error {
	log.Println("DATA received")
	body, err := io.ReadAll(r)
	if err != nil {
		log.Println("Failed reading body:", err)
		return err
	}
	s.data.Write(body)

	msg, err := mail.ReadMessage(strings.NewReader(s.data.String()))
	if err != nil {
		log.Println("error parsing message:", err)
		return nil
	}

	subject := msg.Header.Get("Subject")
	from := msg.Header.Get("From")

	for _, recipient := range s.to {
		parts := strings.Split(recipient, "@")
		if len(parts) < 1 {
			continue
		}
		inboxID := strings.ToLower(strings.TrimSpace(parts[0]))
		key := fmt.Sprintf("inbox:%s:messages", inboxID)

		bodyBytes, _ := io.ReadAll(msg.Body)
		body := string(bodyBytes)

		payload := map[string]interface{}{
			"from":       from,
			"subject":    subject,
			"body":       body,
			"created_at": time.Now().Unix(),
		}

		jsonPayload, _ := json.Marshal(payload)

		redis.Rdb.RPush(context.Background(), key, jsonPayload)
		redis.Rdb.Expire(context.Background(), key, handlers.InboxTTL)

		channel := fmt.Sprintf("inbox:%s", inboxID)
		redis.Rdb.Publish(context.Background(), channel, jsonPayload)
		log.Printf("Received mail for inbox: %s from: %s\n", inboxID, from)
	}
	return nil
}

func (b *Backend) NewSession(c *smtp.Conn) (smtp.Session, error) {
	log.Println("New session created")
	return &Session{}, nil
}

func (s *Session) Reset() {}
func (s *Session) Logout() error {
	return nil
}

func StartSMTPServer() {
	be := &Backend{}
	server := smtp.NewServer(be)

	server.Addr = "0.0.0.0:2525"
	server.Domain = "localhost"
	server.ReadTimeout = 10 * time.Second
	server.WriteTimeout = 10 * time.Second
	server.AllowInsecureAuth = true

	log.Println("Starting SMTP server on", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Failed to start SMTP server:", err)
	}
}
