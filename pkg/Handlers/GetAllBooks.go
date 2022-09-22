package Handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	db "github.com/Golang_assingment/pkg/db"
	user "github.com/Golang_assingment/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllBooks(w http.ResponseWriter, r *http.Request) {

	collection := db.Client()
	var results []*user.User
	cur, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {
		
		// create a value into which the single document can be decoded
		var elem user.User
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())
	w.Header().Add("Content-type","application/json")
	json.NewEncoder(w).Encode(results)
}