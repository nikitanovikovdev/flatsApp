package main

import (
	"flatApp/pkg/platform/repository"
	"log"
)

func main() {
	_, err := repository.NewPostgresDB(repository.Config{
		User:     "postgres",
		Password: "1234567",
		Hostname: "localhost",
		Port:     "5432",
		Database: "postgres",
	})
	if err != nil {
		log.Fatal(err)
	}

	//repo := flats.NewRepository(db)
	//service := flats.NewService(repo)

	//server := flatApp.NewServer(service)

}
