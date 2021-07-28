package internal

import (
	"context"
	"net/http"
)

type Server struct {
	srv http.Server
}

func NewServer(host, port string, h http.Handler) *Server {
	return &Server{
		srv: http.Server{
			Addr:    host + ":" + port,
			Handler: h,
		},
	}
}

func (s *Server) Run() error {
	return s.srv.ListenAndServe()
}

func(s *Server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}