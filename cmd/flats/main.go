package main

import (
	"flatApp"
	"flatApp/pkg/flats"
	"flatApp/pkg/platform/repository"
	_ "github.com/jackc/pgx/v4"
	"github.com/spf13/viper"
	"log"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	db, err := repository.NewPostgresDB(repository.Config{
		User:     viper.GetString("db.user"),
		Password: viper.GetString("db.password"),
		Hostname: viper.GetString("db.hostname"),
		Port:     viper.GetString("db.port"),
		Database: viper.GetString("db.database"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		log.Fatal(err)
	}

	repo := flats.NewRepository(db)
	service := flats.NewService(repo)
	handler := flats.NewHandler(service)

	server := flatApp.NewServer(viper.GetString("server.port"), flats.NewRouter(handler))
	if err = server.Run(); err != nil {
		log.Fatal(err)
	}

}
