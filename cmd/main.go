package main

import (
	"github.com/nikitanovikovdev/flatsApp-flats/internal"
	"github.com/nikitanovikovdev/flatsApp-flats/pkg/flats"
	"github.com/nikitanovikovdev/flatsApp-flats/pkg/platform/repository"
	"github.com/spf13/viper"

	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
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
		Schema:   viper.GetString("db.schema"),
	})

	if err != nil {
		log.Fatal(err)
	}

	repo := flats.NewRepository(db)
	service := flats.NewService(repo)
	handler := flats.NewHandler(service)

	server := internal.NewServer(viper.GetString("server.host"), viper.GetString("server.port"), flats.NewRouter(handler))

	go func() {
		if err = server.Run(); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<- quit

	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatalf("error occured on server shutting down")
	}

	if err := db.Close(); err != nil {
		log.Fatalf("error occured on db connection close")
	}
}

func initConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	return viper.ReadInConfig()
}
