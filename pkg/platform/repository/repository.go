package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
)

type Config struct {
	User     string
	Password string
	Hostname string
	Port     string
	DBName   string
}

func NewPostgresDB(c *Config) (*sql.DB, error) {
	conn := fmt.Sprintf(
		"postgres://%v:%v@%v:%v/%v?sslmode=disable",
		c.User,
		c.Password,
		c.Hostname,
		c.Port,
		c.DBName,
	)

	db, err := sql.Open("pgx", conn)
	if err != nil {
		return &sql.DB{}, err
	}

	return db, nil
}
