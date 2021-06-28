package flatApp

import "net/http"

type Server struct {
	srv http.Server
}

func NewServer(port string) *Server {
	return &Server{
		srv: http.Server{
			Addr: ":" + port,
		},
	}
}
