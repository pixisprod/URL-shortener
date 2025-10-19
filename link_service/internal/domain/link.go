package domain

import "time"

type Link struct {
	ID        int       `db:"id"`
	Hash      string    `db:"hash"`
	URL       string    `db:"url"`
	ExpiresAt time.Time `db:"expires_at"`
}
