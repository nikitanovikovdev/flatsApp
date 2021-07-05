package internal

import (
	"log"
	"net/http"
)

type Server struct {
	srv http.Server
}

func NewServer(port string, h http.Handler) *Server {
	return &Server{
		srv: http.Server{
			Addr:    ":" + port,
			Handler: h,
		},
	}
}

func (s *Server) Run() error {
	if err := s.srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

	return nil
}
