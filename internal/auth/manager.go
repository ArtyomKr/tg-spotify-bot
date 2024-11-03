package auth

import (
	"errors"
	"telegram-bot/internal/spotify"
	"telegram-bot/internal/storage"
	"time"
)

type Manager struct {
	storage    storage.UserStorage
	spotifyAPI *spotify.Client
}

func NewManager(storage storage.UserStorage, spotifyAPI *spotify.Client) *Manager {
	return &Manager{storage, spotifyAPI}
}

func (m *Manager) GetToken(userID string) (string, error) {
	userData, ok := m.storage.Get(userID)
	if !ok {
		return "", errors.New("not authenticated")
	}

	if userData.RefreshToken == "" {
		return m.getInitialToken(userID, userData.Code)
	}
	if m.needsRefresh(userData) {
		return m.refreshToken(userID, userData)
	}

	return userData.AccessToken, nil
}

func (m *Manager) getInitialToken(userID string, code string) (string, error) {
	tokenData, err := m.spotifyAPI.GetToken(code)
	if err != nil {
		return "", err
	}

	m.storage.Set(userID, storage.UserData{
		Code:         code,
		AccessToken:  tokenData.AccessToken,
		RefreshToken: tokenData.RefreshToken,
		ExpiresIn:    tokenData.ExpiresIn,
		ExpiresAt:    time.Now().Add(time.Second * time.Duration(tokenData.ExpiresIn)),
	})

	return tokenData.AccessToken, nil
}

func (m *Manager) refreshToken(userID string, userData storage.UserData) (string, error) {
	tokenData, err := m.spotifyAPI.RefreshToken(userData.RefreshToken)
	if err != nil {
		return "", err
	}

	userData.AccessToken = tokenData.AccessToken
	userData.RefreshToken = tokenData.RefreshToken
	userData.ExpiresIn = tokenData.ExpiresIn
	userData.ExpiresAt = time.Now().Add(time.Second * time.Duration(tokenData.ExpiresIn))

	m.storage.Set(userID, userData)
	return tokenData.AccessToken, nil
}

func (m *Manager) needsRefresh(userData storage.UserData) bool {
	return time.Until(userData.ExpiresAt) < time.Minute
}
