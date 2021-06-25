package repository

import (
	"database/sql"
	"fmt"
)

type Config struct {
	User     string
	Password string
	Hostname string
	Port     string
	Database string
}

func NewPostgresDB(c Config) (*sql.DB, error) {
	conn := fmt.Sprintf(
		"postgres://%v:%v@%v:%v/%v?sslmode=disable",
		c.User,
		c.Password,
		c.Hostname,
		c.Port,
		c.Database)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		return &sql.DB{}, err
	}

	return db, nil
}
