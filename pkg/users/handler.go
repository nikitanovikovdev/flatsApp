package users

import (
	"encoding/json"
	"flatApp/pkg/platform/response"
	"io/ioutil"
	"net/http"
)

type Handler struct {
	s *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{
		s: s,
	}
}

func (h *Handler) SignUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			response.UserError(w, err)
			return
		}

		token, err := h.s.CreateUser(r.Context(), body)
		if err != nil {
			response.DevError(w, err)
			return
		}

		message, err := json.Marshal(token)
		if err != nil {
			response.DevError(w, err)
			return
		}

		response.CreateWithMessage(w, message)
	}
}

func (h *Handler) SignIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			response.UserError(w, err)
			return
		}

		token, err := h.s.GenerateToken(r.Context(), body)
		if err != nil {
			response.UserError(w, err)
			return
		}

		message, err := json.Marshal(token)
		if err != nil {
			response.DevError(w, err)
			return
		}

		response.CreateWithMessage(w, message)
	}
}
