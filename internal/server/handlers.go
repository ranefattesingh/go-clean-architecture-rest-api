package server

import (
	"github.com/ranefattesingh/go-clean-architecture-rest-api/internal/auth/delivery/http"
	"github.com/ranefattesingh/go-clean-architecture-rest-api/internal/auth/mock"

	"github.com/gorilla/mux"
)

func (s *Server) MapHandlers(r *mux.Router) error {

	v1 := r.PathPrefix("/api/v1").Subrouter()

	// Init Repositories
	ar := mock.NewMockRepository()

	// Init Handlers
	ah := http.NewAuthHandler(ar)

	// heahthRoute := v1.PathPrefix("/health").Subrouter()
	auth := v1.PathPrefix("/auth").Subrouter()
	// newsRoute := v1.PathPrefix("/news").Subrouter()
	// commentsRoute := v1.PathPrefix("/comments").Subrouter()

	http.MapAuthRoutes(auth, ah)
	return nil
}
