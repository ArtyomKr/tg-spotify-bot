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
	Set(userID string, data UserData) error
	Get(userID string) (UserData, bool)
	Delete(userID string) error
}

type FileStorage struct {
	mu       sync.Mutex
	data     map[string]UserData
	filepath string
}
