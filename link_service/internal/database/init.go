package database

import (
	"fmt"
	"log"
	"time"

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
	timeout int,
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
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	for {
		err := db.Ping()
		if err == nil {
			log.Println("[db] ✅ Connection set")
			break
		}

		log.Printf("[db] ❌ Connection attempt failed\n⏳ Retrying in %d seconds..\n", timeout)
		time.Sleep(time.Duration(timeout) * time.Second)
	}
	return db, nil
}
