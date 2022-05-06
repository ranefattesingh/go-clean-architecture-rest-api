package interfaces

import "net/http"

type HandlersInterface interface {
	Register() http.HandlerFunc
	Login() http.HandlerFunc
	Logout() http.HandlerFunc
	Update() http.HandlerFunc
	Delete() http.HandlerFunc
	GetUserByID() http.HandlerFunc
	FindByName() http.HandlerFunc
	GetUsers() http.HandlerFunc
	GetMe() http.HandlerFunc
	UploadAvatar() http.HandlerFunc
	GetCSRFToken() http.HandlerFunc
}
