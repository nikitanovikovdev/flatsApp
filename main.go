package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)
type Flat struct {
	Id string `json:"id"`
	City *City `json:"city"`
	Street string	`json:"street"`
	HouseNumber string	`json:"house_number"`
	RoomNumber int	`json:"room_number"`
	Description string	`json:"description"`
}

type City struct {
	Country string	`json:"country"`
	Name string	`json:"name"`
}

var DB = make(map[string]Flat)

func main() {
	r := mux.NewRouter()

	DB["1"] = Flat{"1",&City{"Belarus","Minsk"},"Lenina","12A", 23, "Хорошая квартира"}
	DB["2"] = Flat{"2",&City{"Russia","Moscow"},"Tolstogo","23", 123, "Квартира возле кремля"}


	r.HandleFunc("/flats", createFlat).Methods("POST")
	r.HandleFunc("/flats", getFlats).Methods("GET")
	r.HandleFunc("/flats/{id}", getFlat).Methods("GET")
	r.HandleFunc("/flats/{id}", updateFlat).Methods("PUT")
	r.HandleFunc("/flats/{id}", deleteFlat).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))

}

func createFlat(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var flat Flat
	_ = json.NewDecoder(r.Body).Decode(&flat)
	flat.Id = strconv.Itoa(rand.Intn(100))
	DB[flat.Id] = flat
	_ = json.NewEncoder(w).Encode(flat)
}



func getFlats(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(DB); if err != nil {
		log.Fatal(err)
	}
}

func getFlat(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for _,item := range DB {
		if item.Id == params["id"] {
			err := json.NewEncoder(w).Encode(item); if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func updateFlat(w http.ResponseWriter, r *http.Request) {

}

func deleteFlat(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range DB {
		if item.Id == params["id"] {
			delete(DB,item.Id)
		}
	}
}