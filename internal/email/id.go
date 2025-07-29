package email

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateInboxId() string {
	bytes := make([]byte, 4)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}
