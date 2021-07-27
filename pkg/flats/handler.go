package flats

import (
	"encoding/json"
	"flatApp/pkg/platform/middlewear"
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
		token := r.Header["Token"][0]

		username, _ := middlewear.ParseToken(token)

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			response.UserError(w, err)
			return
		}

		flat, err := h.service.Create(r.Context(), body, username)
		if err != nil {
			response.UserError(w, err)
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
		token := r.Header["Token"][0]

		username, _ := middlewear.ParseToken(token)

		flat, err := h.service.Read(r.Context(), username)
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
		}

		response.OkWithMessage(w, message)
	}
}

func (h *Handler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]

		token := r.Header["Token"][0]
		username, _ := middlewear.ParseToken(token)

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			response.UserError(w, err)
			return
		}

		if err := h.service.Update(r.Context(), id, body, username); err != nil {
			response.DevError(w, err)
			return
		}

		response.Ok(w)
	}
}

func (h *Handler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]

		token := r.Header["Token"][0]
		username, _ := middlewear.ParseToken(token)

		if err := h.service.Delete(r.Context(), id, username); err != nil {
			response.UserError(w, err)
			return
		}

		response.Ok(w)
	}
}

