package storage

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func NewSqlLiteStorage(path string) (*SQLiteUserStorage, error) {
	db, err := sql.Open("sqlite3", path)

	if err != nil {
		return nil, fmt.Errorf("could not open database %w", err)
	}

	createUserTableSql := `
	CREATE TABLE IF NOT EXISTS user_tokens (
		user_id TEXT NOT NULL PRIMARY KEY,
		code TEXT,
		access_token TEXT,
		expires_in INTEGER,
		refresh_token TEXT,
		expires_at DATETIME
	);`

	if _, err := db.Exec(createUserTableSql); err != nil {
		return nil, fmt.Errorf("could not create table: %w", err)
	}

	return &SQLiteUserStorage{db: db}, nil
}

func (s *SQLiteUserStorage) Set(userID string, data UserData) error {
	query := `
	INSERT OR REPLACE INTO user_tokens 
	(user_id, code, access_token, expires_in, refresh_token, expires_at) 
	VALUES (?, ?, ?, ?, ?, ?);`

	if _, err := s.db.Exec(query, userID, data.Code, data.AccessToken, data.ExpiresIn, data.RefreshToken, data.ExpiresAt); err != nil {
		return fmt.Errorf("could not set user data for %s: %w", userID, err)
	}
	return nil
}

func (s *SQLiteUserStorage) Get(userID string) (UserData, bool) {
	query := `
	SELECT code, access_token, expires_in, refresh_token, expires_at
	FROM user_tokens WHERE user_id = ?;`

	var userData UserData
	row := s.db.QueryRow(query, userID)

	err := row.Scan(&userData.Code, &userData.AccessToken, &userData.ExpiresIn, &userData.RefreshToken, &userData.ExpiresAt)

	if errors.Is(err, sql.ErrNoRows) {
		return UserData{}, false
	}

	if err != nil {
		return UserData{}, false
	}

	return userData, true
}

func (s *SQLiteUserStorage) Delete(userID string) error {
	query := `DELETE FROM user_tokens WHERE user_id = ?;`

	if _, err := s.db.Exec(query, userID); err != nil {
		return fmt.Errorf("could not delete data for user %s", userID)
	}
	return nil
}

func (s *SQLiteUserStorage) Close() {
	s.db.Close()
}
