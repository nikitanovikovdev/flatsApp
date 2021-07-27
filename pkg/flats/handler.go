package flats

import (
	"encoding/json"
	"flatApp/pkg/platform/response"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
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
			response.UserError(w, err)
			return
		}

		flat, err := h.service.Create(r.Context(), body)
		if err != nil {
			response.DevError(w, err)
			return
		}

		message, err := json.Marshal(flat)
		if err != nil {
			response.DevError(w, err)
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
			response.UserError(w, err)
			return
		}

		message, err := json.Marshal(flat)
		if err != nil {
			response.DevError(w, err)
			return
		}

		response.OkWithMessage(w, message)
	}
}

func (h *Handler) ReadAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		flat, err := h.service.ReadAll(r.Context())
		if err != nil {
			response.UserError(w, err)
			return
		}

		message, err := json.Marshal(flat)
		if err != nil {
			response.DevError(w, err)
			return
		}

		response.OkWithMessage(w, message)
	}
}

func (h *Handler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vin := mux.Vars(r)["id"]

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			response.UserError(w, err)
			return
		}

		if err := h.service.Update(r.Context(), vin, body); err != nil {

			response.DevError(w, err)
			return
		}

		response.Ok(w)
	}
}

func (h *Handler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]

		if err := h.service.Delete(r.Context(), id); err != nil {
			response.UserError(w, err)
			return
		}

		response.Ok(w)
	}
}
