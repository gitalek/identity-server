package handler

import (
	"net/http"
	"server/models/authenticator"
)

type handler struct {
	authenticator authenticator.Authenticable
}

func (h *handler) SignUp(w http.ResponseWriter, r *http.Request) {}

func (h *handler) SignIn(w http.ResponseWriter, r *http.Request) {}

func NewHandler(authenticator authenticator.Authenticable) *handler {
	return &handler{
		authenticator: authenticator,
	}
}