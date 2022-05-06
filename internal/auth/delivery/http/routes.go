package http

import (
	"net/http"

	"github.com/ranefattesingh/go-clean-architecture-rest-api/internal/auth/interfaces"

	"github.com/gorilla/mux"
)

func MapAuthRoutes(r *mux.Router, h interfaces.HandlersInterface) {

	// POST routes
	r.HandleFunc("/register", h.Register()).Methods(http.MethodPost)
	r.HandleFunc("/login", h.Login()).Methods(http.MethodPost)
	r.HandleFunc("/logout", h.Logout()).Methods(http.MethodPost)

	// GET routes
	get := r.Methods(http.MethodGet)
	get.Path("/find").HandlerFunc(h.FindByName())
	get.Path("/all").HandlerFunc(h.GetUsers())
	get.Path("/{id:[0-9]+}").HandlerFunc(h.GetUserByID())

	// Middleware routes

}
