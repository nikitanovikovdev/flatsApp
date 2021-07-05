package flats

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return
		}

		ids, err := h.service.Create(r.Context(), body)
		if err != nil {
			return
		}

		message, err := json.Marshal(ids)
		if err != nil {
			return
		}

		if _, err := w.Write(message); err != nil {
			log.Println(err.Error())
		}
	}
}

func (h *Handler) Read() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]

		flat, err := h.service.Read(r.Context(), id)
		if err != nil {
			fmt.Println("handler read error")
			return
		}

		message, err := json.Marshal(flat)
		if err != nil {
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if _, err := w.Write(message); err != nil {
			log.Println(err.Error())
		}
	}
}

func (h *Handler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return
		}

		vin := mux.Vars(r)["id"]

		if err := h.service.Update(r.Context(), vin, body); err != nil {
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}

func (h *Handler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]

		if err := h.service.Delete(r.Context(), id); err != nil {
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}
