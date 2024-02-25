package pkg

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

type Server struct {
	port    uint32
}

func NewServer(ctx context.Context, port uint32) *Server {
	return &Server{
		port:    port,
	}
}

func (s *Server) Run() error {
	http.HandleFunc("/", handleReq)
	return http.ListenAndServe(fmt.Sprintf(":%d", s.port), nil)
}

func handleReq(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Hello there!")
}
