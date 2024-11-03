package storage

import (
	"sync"
	"time"
)

type UserData struct {
	Code         string
	AccessToken  string
	ExpiresIn    int
	RefreshToken string
	ExpiresAt    time.Time
}

type UserStorage interface {
	Set(userID string, data UserData)
	Get(userID string) (UserData, bool)
	Delete(userID string)
}

type MemoryStorage struct {
	mu   sync.Mutex
	data map[string]UserData
}
