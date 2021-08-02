package flats

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/nikitanovikovdev/flatsApp-flats/pkg/middlewear"
	"github.com/nikitanovikovdev/flatsApp-flats/pkg/platform/response"
	"github.com/nikitanovikovdev/flatsApp-flats/pkg/platform/user"
	authorizations "github.com/nikitanovikovdev/flatsApp-users/proto"

	"google.golang.org/grpc"
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
		cookie, err  := r.Cookie("token")
		if err != nil {
			response.DevError(w, err)
			return
		}

		token := cookie.Value

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
		cookie, err  := r.Cookie("token")
		if err != nil {
			response.DevError(w, err)
			return
		}

		token := cookie.Value

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

		cookie, err  := r.Cookie("token")
		if err != nil {
			response.DevError(w, err)
			return
		}

		token := cookie.Value

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

		cookie, err  := r.Cookie("token")
		if err != nil {
			response.DevError(w, err)
			return
		}

		token := cookie.Value

		username, _ := middlewear.ParseToken(token)

		if err := h.service.Delete(r.Context(), id, username); err != nil {
			response.UserError(w, err)
			return
		}

		response.Ok(w)
	}
}

func (h *Handler) AuthorizationHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user user.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			response.UserError(w, err)
			return
		}

		conn, err := grpc.Dial(":8040", grpc.WithInsecure())
		if err != nil {
			response.DevError(w, err)
			return
		}

		c := authorizations.NewAuthClient(conn)

		tkn, err := c.Authorize(r.Context(), &authorizations.RequestData{Username: user.Username, Password: user.Password})
		if err != nil {
			response.DevError(w, err)
			return
		}

		token := tkn.GetToken()
		if token == "" {
			response.InvalidToken(w)
			return
		}

		cookie, err := h.service.Authorize(token)
		if err != nil {
			response.DevError(w, err)
			return
		}

		http.SetCookie(w, &cookie)
		response.OkWithMessage(w, []byte(token))
	}
}

func (h *Handler) RegistrationHandler() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		var user user.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			response.UserError(w, err)
			return
		}

		conn, err := grpc.Dial(":8040", grpc.WithInsecure())
		if err != nil {
			response.DevError(w, err)
			return
		}

		c := authorizations.NewAuthClient(conn)

		idRes, err := c.Registr(r.Context(), &authorizations.RegistrData{Username: user.Username, Password: user.Password})
		if err != nil {
			response.UserError(w,err)
			return
		}
		id := idRes.GetId()

		if id == "" {
			response.RegistrError(w)
			return
		}

		response.OkWithMessage(w,[]byte(id))
	}
}