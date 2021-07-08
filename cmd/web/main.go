package main

import (
	"flatApp/internal"
	"flatApp/pkg/flats"
	"flatApp/pkg/platform/repository"
	"github.com/pkg/errors"
	"log"

	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing config: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(&repository.Config{
		User:     viper.GetString("db.user"),
		Password: viper.GetString("db.password"),
		Hostname: viper.GetString("db.hostname"),
		Port:     viper.GetString("db.port"),
		DBName:   viper.GetString("db.dbname"),
	})

	if err != nil {
		log.Fatal(err)
	}

	repo := flats.NewRepository(db)
	service := flats.NewService(repo)
	handler := flats.NewHandler(service)

	server := internal.NewServer(viper.GetString("server.host"), viper.GetString("server.port"), flats.NewRouter(handler))
	if err = server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	return errors.Wrap(viper.ReadInConfig(), "error reading config")
}
