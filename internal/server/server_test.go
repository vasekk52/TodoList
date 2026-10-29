package server

import (
	"context"
	"net/http"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	mux := http.NewServeMux()
	srv := New("8080", mux)

	if srv == nil || srv.server == nil {
		t.Fatal("nil server")
	}
	if srv.server.Addr != ":8080" {
		t.Errorf("got %v, want :8080", srv.server.Addr)
	}
}

func  TestServerGracefulShutdown(t *testing.T) {
	mux := http.NewServeMux()
	srv := New("0", mux)

	errCh := make(chan error, 1)
	go func(){
	  	errCh <- srv.StartServer()
	}()

	time.Sleep(100 *time.Millisecond)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := srv.server.Shutdown(ctx); err != nil {
		t.Fatalf("failed to shutdown server %v", err)
	}
}

func TestServerErrorPort(t *testing.T){
	mux := http.NewServeMux()
	srv := New("testport", mux)

	err := srv.StartServer()
	if err == nil {
		t.Fatal("expected error port, got nil")
	}
}