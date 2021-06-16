package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
	"log"
	"net/http"
)

type Flat struct {
	Id          int    `json:"id"`
	Street      string `json:"street"`
	HouseNumber string `json:"house_number"`
	RoomNumber  int    `json:"room_number"`
	Description string `json:"description"`
	City        City   `json:"city"`
}

type City struct {
	Id      int    `json:"id"`
	Country string `json:"country"`
	Name    string `json:"name"`
}

var conn *pgx.Conn

func main() {
	var err error

	url := fmt.Sprintf(
		"postgres://%v:%v@%v:%v/%v?sslmode=disable",
		"postgres",
		"123456",
		"localhost",
		"5432",
		"postgres")

	conn, err = pgx.Connect(context.Background(), url)
	if err != nil {
		fmt.Println(err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/flats", createFlat).Methods("POST")
	r.HandleFunc("/flats", getFlats).Methods("GET")
	r.HandleFunc("/flats/{id}", getFlat).Methods("GET")
	r.HandleFunc("/flats/{id}", updateFlat).Methods("PUT")
	r.HandleFunc("/flats/{id}", deleteFlat).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))

}

func createFlat(w http.ResponseWriter, r *http.Request) {
	var f Flat

	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&f)
	if err != nil {
		log.Fatal(err)
	}

	query := "INSERT INTO flats (street,house_number,room_number,description,city_id) VALUES ($1,$2,$3,$4,$5)"

	insert, err := conn.Query(context.Background(), query, f.Street, f.HouseNumber, f.RoomNumber, f.Description, f.City.Id)
	if err != nil {
		log.Fatal(err)
	}
	defer insert.Close()

}

func getFlats(w http.ResponseWriter, r *http.Request) {
	rows, _ := conn.Query(context.Background(), "SELECT * FROM flats")
	defer rows.Close()

	var f Flat
	for rows.Next() {
		err := rows.Scan(
			&f.Id,
			&f.Street,
			&f.HouseNumber,
			&f.RoomNumber,
			&f.Description,
			&f.City.Id)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Printf("%d, %s,%s,%d,%s,%d\n", f.Id, f.Street,f.HouseNumber,f.RoomNumber,f.Description,f.City.Id)

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(f)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func getFlat(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var f Flat

	err := conn.QueryRow(context.Background(), "SELECT * FROM flats WHERE id=$1", params["id"]).Scan(
		&f.Id,
		&f.Street,
		&f.HouseNumber,
		&f.RoomNumber,
		&f.Description,
		&f.City.Id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&f)
	if err != nil {
		log.Fatal(err)
	}
}

func updateFlat(w http.ResponseWriter, r *http.Request) {
	var f Flat

	params := mux.Vars(r)

	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&f)
	if err != nil {
		log.Fatal(err)
	}

	query := "UPDATE flats SET street = $1, house_number = $2, room_number = $3, description = $4, city_id = $5 WHERE id =$6 "

	_, err = conn.Exec(context.Background(), query, f.Street, f.HouseNumber, f.RoomNumber, f.Description, f.City.Id, params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}

func deleteFlat(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	_, err := conn.Exec(context.Background(), "DELETE FROM flats WHERE id=$1", params["id"])
	if err != nil {
		log.Fatal(err)
	}
}
