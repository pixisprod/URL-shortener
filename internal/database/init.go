package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func InitDbPool(
	user string,
	password string,
	host string,
	port int,
	name string,
	sslmode string,
) (*sqlx.DB, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		user,
		password,
		host,
		port,
		name,
		sslmode,
	)
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
