package main

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
	
	//"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	clientsOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientsOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	r := mux.NewRouter()

	r.Handle("/", http.FileServer(http.Dir("./views/"))).Methods("GET")
	r.Handle("/users/{name}", http.HandlerFunc(UserHandler)).Methods("GET")

	r.Handle("/get-token", http.HandlerFunc(GetTokenHandler)).Methods("GET")
	r.Handle("/users/{guid}/tokens", http.HandlerFunc(GenerateTokenHandler)).Methods("POST")

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

func GenerateTokenHandler(w http.ResponseWriter, r *http.Request) {
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
