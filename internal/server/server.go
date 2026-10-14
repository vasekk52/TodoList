package server

import (
	"net/http"
	"time"
)

type Server struct {
	server *http.Server
}

func New(addr string, handler http.Handler) *Server {
	return &Server{
		server: &http.Server{
			Addr: ":" + addr,
			Handler: handler,
			ReadTimeout: 5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout: 30 * time.Second,
		},
	}
}

func (s *Server) StartServer() error{

	if err := s.server.ListenAndServe();  err != nil {
		
		return err 
	}

	return nil
}