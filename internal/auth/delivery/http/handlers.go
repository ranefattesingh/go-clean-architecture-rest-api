package http

import (
	"log"
	"net/http"

	"github.com/ranefattesingh/go-clean-architecture-rest-api/internal/auth/interfaces"
	"github.com/ranefattesingh/go-clean-architecture-rest-api/internal/models"
	"github.com/ranefattesingh/go-clean-architecture-rest-api/pkg/utils"
)

type authHandler struct {
	ar interfaces.RepositoryInterface
}

func NewAuthHandler(ar interfaces.RepositoryInterface) *authHandler {
	return &authHandler{ar}
}

func (h *authHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u := &models.User{}
		if err := utils.FromJSON(r.Body, u); err != nil {
			log.Println("Error")
		}

		log.Println(u)

		h.ar.Register(r.Context(), u)

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Created"))
	}
}

func (h *authHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *authHandler) Logout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *authHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *authHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *authHandler) GetUserByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *authHandler) FindByName() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *authHandler) GetUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *authHandler) GetMe() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *authHandler) UploadAvatar() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *authHandler) GetCSRFToken() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
