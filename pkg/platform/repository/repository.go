package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/stdlib"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	User     string
	Password string
	Hostname string
	Port     string
	DBName   string
	Schema   string
}

func NewPostgresDB(c *Config) (*sql.DB, error) {
	conn := fmt.Sprintf(
		"postgres://%v:%v@%v:%v/%v?sslmode=disable&search_path=%s, public",
		c.User,
		c.Password,
		c.Hostname,
		c.Port,
		c.DBName,
		c.Schema,
	)

	db, err := sql.Open("pgx", conn)
	if err != nil {
		return &sql.DB{}, err
	}

	return db, nil
}

type MongoConfig struct {
	Host     string
	Port     string
}

func NewMongoDB(c *MongoConfig) (*mongo.Client, error) {
	//uri := "mongodb://localhost:27017"
	uri := fmt.Sprintf("mongodb://%v:%v", c.Host, c.Port)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	return client, nil
}