package Handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	db "github.com/Golang_assingment/pkg/db"
	user "github.com/Golang_assingment/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetBook(w http.ResponseWriter,r *http.Request){
	//read dynamic id parameter
	vars:= mux.Vars(r)
	id,_:=strconv.Atoi(vars["id"])
	collection := db.Client()
	log.Print(id)
	var result user.User
	err := collection.FindOne(context.TODO(), bson.D{{Key: "id", Value: id}}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}else{
		w.Header().Add("Content-type","application/json")
		json.NewEncoder(w).Encode(result)
	}

}