package main

import (
	"context"
	"flatApp/internal"
	"flatApp/pkg/platform/repository"
	"flatApp/pkg/users"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatal(err)
	}

	db, err := repository.NewMongoDB(&repository.MongoConfig{
		Host: viper.GetString("mongodb.host"),
		Port: viper.GetString("mongodb.port"),
	})
	if err != nil {
		log.Fatalf("failed to create new mingo repository: %v", err)
	}
	repo := users.NewRepository(db)
	service := users.NewService(repo)
	handler := users.NewHandler(service)
	server := internal.NewServer(viper.GetString("server_user.hostname"), viper.GetString("server_user.port"), users.NewRouter(handler))

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

	if err := db.Disconnect(context.Background()); err != nil {
		log.Fatalf("error occured on db connection close")
	}
}

func initConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	return viper.ReadInConfig()
}
