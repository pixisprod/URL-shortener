package repository

import (
	"github.com/jmoiron/sqlx"
)

type LinkRepository struct {
	Db *sqlx.DB
}

func (lr *LinkRepository) Add(hash string, url string) error {
	query := "INSERT INTO links (hash, url) VALUES ($1, $2);"
	_, err := lr.Db.Exec(query, hash, url)
	if err != nil {
		return err
	}
	return nil
}

func (lr *LinkRepository) GetByHash(hash string) (string, error) {
	var url string
	query := "SELECT url FROM links WHERE hash = $1;"
	err := lr.Db.Get(&url, query, hash)
	if err != nil {
		return "", err
	}
	return url, nil
}
