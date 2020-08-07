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

	type (
		User struct {
			GUID         string
			Name         string
			RefreshToken string
			AccessToken  string
		}
	)

	users := client.Database("server").Collection("users")
	
	alex := User{
		GUID:         "4aabc26c-55f5-4320-8635-33cc02023028",
		Name:         "Alex",
		RefreshToken: "",
		AccessToken:  "",
	}

	peter := User{
		GUID:         "36b1927d-8dcc-4a92-bf71-a75420d1ca25",
		Name:         "Peter",
		RefreshToken: "",
		AccessToken:  "",
	}

	nick := User{
		GUID:         "de6f996a-f8f0-4683-813d-7fb0a7f84656",
		Name:         "Nick",
		RefreshToken: "",
		AccessToken:  "",
	}

	res, err := users.InsertMany(context.TODO(), []interface{}{alex, peter, nick})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted multiple documents: ", res)
	fmt.Println("Inserted ids: ", res.InsertedIDs)

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
