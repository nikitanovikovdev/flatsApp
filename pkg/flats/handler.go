package flats

import (
	"encoding/json"
	"flatApp/pkg/platform/response"
	"io/ioutil"
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
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			response.Bad(w, err)
			return
		}

		ids, err := h.service.Create(r.Context(), body)
		if err != nil {
			response.Bad(w, err)
			return
		}

		message, err := json.Marshal(ids)
		if err != nil {
			response.Bad(w, err)
			return
		}

		response.CreateWithMessage(w, message)
	}
}

func (h *Handler) Read() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]

		flat, err := h.service.Read(r.Context(), id)
		if err != nil {
			response.Bad(w, err)
			return
		}

		message, err := json.Marshal(flat)
		if err != nil {
			response.Bad(w, err)
			return
		}

		response.OkWithMessage(w, message)
	}
}

func (h *Handler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			response.Bad(w, err)
			return
		}

		vin := mux.Vars(r)["id"]

		if err := h.service.Update(r.Context(), vin, body); err != nil {
			response.Bad(w, err)
			return
		}

		response.Ok(w)
	}
}

func (h *Handler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]

		if err := h.service.Delete(r.Context(), id); err != nil {
			return
		}

		response.Ok(w)
	}
}
