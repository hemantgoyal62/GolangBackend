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
	filter := bson.D{{Key: "Id",Value: id}}
	update := bson.D{
		{Key: "$Name",Value: updatedBook.Name},
		{Key: "$Address",Value: updatedBook.Address},
		{Key: "$Desc",Value: updatedBook.Desc},
	}
	updateResult,err := collection.UpdateOne(context.TODO(),filter,update)
	if err!=nil{
		log.Fatal(err)
	} else{
		w.Header().Add("Content-type","application/json")
		json.NewEncoder(w).Encode(updateResult)
	}
	// for index,book := range mocks.Books{
	// 	if book.Id == id {
	// 		book.Title = updatedBook.Title
	// 		book.Author = updatedBook.Author
	// 		book.Desc = updatedBook.Desc
	// 		book.Id = id
	// 		mocks.Books[index] = book
	// 		w.WriteHeader(http.StatusOK)
	// 		w.Header().Add("Content-Type","applicaiton/json")
	// 		json.NewEncoder(w).Encode("updated")
	// 		json.NewEncoder(w).Encode(mocks.Books)
	// 		json.NewEncoder(w).Encode(updatedBook)
	// 		break
	// 	}
	// }
}