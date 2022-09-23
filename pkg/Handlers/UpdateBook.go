package Handlers

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	db "github.com/Golang_assingment/pkg/db"
	user "github.com/Golang_assingment/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateBook(w http.ResponseWriter,r *http.Request){
	
	vars:=mux.Vars(r)
	id,_ := strconv.Atoi(vars["id"])
	

	defer r.Body.Close()
	body,err := ioutil.ReadAll(r.Body)
	if err != nil{
		log.Fatalln(err)
	}

	var updatedBook user.User
	json.Unmarshal(body, &updatedBook)


	collection := db.Client()
	filter := bson.D{{Key: "id",Value: id}}
	update := bson.D{
		{Key: "$set",Value: bson.D{{Key: "Name",Value: updatedBook.Name}}},
		{Key: "$set",Value: bson.D{{Key: "Address",Value: updatedBook.Address}}},
		{Key: "$set",Value: bson.D{{Key: "Desc",Value: updatedBook.Desc}}},
	}
	updateResult,err := collection.UpdateOne(context.TODO(),filter,update)
	if err!=nil{
		log.Fatal(err)
	} else{
		w.Header().Add("Content-type","application/json")
		json.NewEncoder(w).Encode(updateResult)
	}
}