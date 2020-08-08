package app

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"server/models/authenticator"
	"server/models/handler"
)

type App struct {
	server *http.Server
	authenticator authenticator.Authenticable
}

func (app *App) Run(port string) error {
	r := mux.NewRouter()

	r.Handle("/", http.FileServer(http.Dir("./views/"))).Methods("GET")

	h := handler.NewHandler(app.authenticator)
	r.Handle("/sign-up", http.HandlerFunc(h.SignUp)).Methods("POST")
	r.Handle("/sign-in", http.HandlerFunc(h.SignIn)).Methods("POST")

	http.ListenAndServe(":3333", handlers.LoggingHandler(os.Stdout, r))

	return nil
}