package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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
	errCh := make(chan error, 1)

	go func() {
	if err := s.server.ListenAndServe();  err != nil && !errors.Is(err, http.ErrServerClosed) {
		errCh <-err 
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-errCh:
		return err
	
	case sig := <-stop:
		fmt.Printf("shutdown signal received: %s\n", sig)

		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		if err := s.server.Shutdown(ctx); err != nil {
		  	return err
		}
		fmt.Println("server stopped")
	}
	return nil
}