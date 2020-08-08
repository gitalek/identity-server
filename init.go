package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func initDB() *mongo.Database {
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
	return client.Database("server")
}

//alex := User{
//	GUID:         "4aabc26c-55f5-4320-8635-33cc02023028",
//	Name:         "Alex",
//	RefreshToken: "",
//	AccessToken:  "",
//}
//
//peter := User{
//	GUID:         "36b1927d-8dcc-4a92-bf71-a75420d1ca25",
//	Name:         "Peter",
//	RefreshToken: "",
//	AccessToken:  "",
//}
//
//nick := User{
//	GUID:         "de6f996a-f8f0-4683-813d-7fb0a7f84656",
//	Name:         "Nick",
//	RefreshToken: "",
//	AccessToken:  "",
//}

//res, err := users.InsertMany(context.TODO(), []interface{}{alex, peter, nick})
//if err != nil {
//	log.Fatal(err)
//}
//fmt.Println("Inserted multiple documents: ", res)
//fmt.Println("Inserted ids: ", res.InsertedIDs)
