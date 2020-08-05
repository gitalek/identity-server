package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"time"
)

func main() {
	r := mux.NewRouter()

	r.Handle("/get-token", http.HandlerFunc(GetTokenHandler)).Methods("GET")
	r.Handle("/", http.FileServer(http.Dir("./views/"))).Methods("GET")
	r.Handle("/users/{name}", http.HandlerFunc(UserHandler)).Methods("GET")

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/",
		http.FileServer(http.Dir("./static/")))).Methods("GET")

	http.ListenAndServe(":3333", handlers.LoggingHandler(os.Stdout, r))
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Your name is %v\n", vars["name"])
}

var mySigningKey = []byte("secret")

func GetTokenHandler(w http.ResponseWriter, r *http.Request) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"admin": true,
			"name": "Alex O",
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, _ := token.SignedString(mySigningKey)
	fmt.Fprintf(w, tokenString)
}