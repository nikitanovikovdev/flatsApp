package internal

import (
	"github.com/pkg/errors"
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
	if err := s.srv.ListenAndServe(); err != nil {
		return errors.Wrap(err, "listening or serving failed")
	}

	return nil
}
