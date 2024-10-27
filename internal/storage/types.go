package storage

import "sync"

type UserData struct {
	Code  string
	Token string
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
