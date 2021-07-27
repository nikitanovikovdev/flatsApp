package users

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	handler *Handler
}

func NewRouter(h *Handler) http.Handler {
	router := &Router{
		handler: h,
	}

	return router.initRoutes()
}

func (r *Router) initRoutes() http.Handler {
	m := mux.NewRouter()

	m.HandleFunc("/auth/sign-up", r.handler.SignUp()).Methods(http.MethodPost)
	m.HandleFunc("/auth/sign-in", r.handler.SignIn()).Methods(http.MethodPost)

	return m
}
