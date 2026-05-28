package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Client struct {
	db *sql.DB
}

func NewClient(pathToDB string) (Client, error) {
	db, err := sql.Open("sqlite3", pathToDB)
	if err != nil {
		return Client{}, fmt.Errorf("couldn't open database: %w", err)
	}
	err = db.Ping()
	if err != nil {
		return Client{}, fmt.Errorf("couldn't ping database: %w", err)
	}
	return Client{db: db}, nil
}

func (c Client) CreateTables() error {
	_, err := c.db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id TEXT PRIMARY KEY,
			email TEXT UNIQUE NOT NULL,
			hashed_password TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);

		CREATE TABLE IF NOT EXISTS tracks (
			id TEXT PRIMARY KEY,
			title TEXT NOT NULL,
			album TEXT,
			year INTEGER,
			duration_seconds INTEGER,
			s3_bucket TEXT,
			s3_key TEXT,
			track_url TEXT,
			artwork_url TEXT,
			is_preview INTEGER DEFAULT 0,
			user_id TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);

		CREATE TABLE IF NOT EXISTS videos (
			id TEXT PRIMARY KEY,
			title TEXT NOT NULL,
			description TEXT,
			s3_bucket TEXT,
			s3_key TEXT,
			video_url TEXT,
			thumbnail_url TEXT,
			aspect_ratio TEXT,
			is_preview INTEGER DEFAULT 0,
			user_id TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);

		CREATE TABLE IF NOT EXISTS refresh_tokens (
			token TEXT PRIMARY KEY,
			user_id TEXT NOT NULL,
			expires_at TIMESTAMP NOT NULL,
			revoked_at TIMESTAMP
		);
	`)
	return err
}
