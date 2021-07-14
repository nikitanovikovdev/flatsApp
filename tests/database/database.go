package database

import (
	"database/sql"
	"flatApp/pkg/flats"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/jackc/pgx/stdlib"
	"log"
)

func CreateTestFlatsRepository(schema string) *flats.RepositorySQL {
	psqlConn := fmt.Sprintf("host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable search_path=%s", schema)

	db, err := sql.Open("pgx", psqlConn)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	_, err = db.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", schema))
	if err != nil {
		log.Fatal(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	//fsrc, err := (&file.File{}).Open("../../../migrations")
	//if err != nil {
	//	log.Fatal(err)
	//}

	//m, err := migrate.NewWithInstance(
	//	"file", fsrc,
	//	"postgres", driver)
	//if err != nil {
	//	log.Fatal(err)
	//}

	m, err := migrate.NewWithDatabaseInstance(
		"../../../migrations",
		"postgres",driver)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Steps(6); err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
	defer func(){
		if err := m.Drop(); err != nil {
			log.Fatal(err)
		}
	}()

	repo := flats.NewRepository(db)

	return repo
}