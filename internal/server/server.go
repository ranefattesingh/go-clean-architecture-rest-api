package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

// Server struct
type Server struct {
	r *mux.Router
}

// NewServer New Server constructor
func NewServer() *Server {
	return &Server{r: mux.NewRouter()}
}

func (s *Server) Run() error {

	s.MapHandlers(s.r)

	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: s.r,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	msg := <-quit

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	log.Println("Server terminated successfully because of ", msg.String())
	return server.Shutdown(ctx)
}
