package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4"
)

type Config struct {
	User     string
	Password string
	Hostname string
	Port     string
	Database string
	SSLMode  string
}

func NewPostgresDB(c Config) (*sql.DB, error) {
	conn := fmt.Sprintf(
		"postgres://%v:%v@%v:%v/%v?sslmode=%v",
		c.User,
		c.Password,
		c.Hostname,
		c.Port,
		c.Database,
		c.SSLMode)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		return &sql.DB{}, err
	}

	return db, nil
}
