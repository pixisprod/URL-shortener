package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/pixisprod/URL-shortener/internal/domain"
)

type LinkRepository struct {
	Db *sqlx.DB
}

func (lr *LinkRepository) Add(link domain.Link) error {
	query := `
		INSERT INTO links (hash, url, expires_at) 
		VALUES ($1, $2, $3)
	`
	_, err := lr.Db.Exec(query, link.Hash, link.URL, link.ExpiresAt)
	if err != nil {
		return err
	}
	return nil
}

func (lr *LinkRepository) GetByHash(hash string) (domain.Link, error) {
	var link domain.Link
	query := `
		SELECT id, hash, url, expires_at FROM links
		WHERE hash = $1;
	`
	err := lr.Db.Get(&link, query, hash)
	if err != nil {
		return domain.Link{}, err
	}
	return link, nil
}
