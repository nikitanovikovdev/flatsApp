package flats

import (
	"github.com/gorilla/mux"
	"net/http"
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

	m.HandleFunc("/flats", r.handler.Create()).Methods("POST")
	m.HandleFunc("/flats/{id}", r.handler.Read()).Methods("GET")
	m.HandleFunc("/flats/{id}", r.handler.Update()).Methods("PUT")
	m.HandleFunc("/flats/{id}", r.handler.Delete()).Methods("DELETE")

	return m
}
