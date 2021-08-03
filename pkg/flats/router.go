package flats

import (
	"github.com/gorilla/mux"
	"github.com/nikitanovikovdev/flatsApp-flats/pkg/middlewear"
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

	m.Handle("/flats", middlewear.IsAuthorized(r.handler.Create())).Methods(http.MethodPost)
	m.Handle("/flats/", middlewear.IsAuthorized(r.handler.Read())).Methods(http.MethodGet)
	m.Handle("/flats", middlewear.IsAuthorized(r.handler.ReadAll())).Methods(http.MethodGet)
	m.Handle("/flats/{id}", middlewear.IsAuthorized(r.handler.Update())).Methods(http.MethodPut)
	m.Handle("/flats/{id}", middlewear.IsAuthorized(r.handler.Delete())).Methods(http.MethodDelete)
	//m.Handle("/authorization", r.handler.AuthorizationHandler()).Methods(http.MethodPost)
	//m.Handle("/registration", r.handler.RegistrationHandler()).Methods(http.MethodPost)

	return m
}