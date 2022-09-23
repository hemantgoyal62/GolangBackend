package Handlers

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	db "github.com/Golang_assingment/pkg/db"
	"github.com/Golang_assingment/pkg/models"
)

func AddBook(w http.ResponseWriter,r *http.Request) {
 //Read to request Body
 defer r.Body.Close()
 body, err:=ioutil.ReadAll(r.Body)
 if err !=nil{
	log.Fatal(err)
 }
 var user models.User
 json.Unmarshal(body,&user)
 user.CreatedAt = time.Now()
 collection := db.Client()
 insertresult,err := collection.InsertOne(context.TODO(),user)
 if err!=nil{
	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(err)
	json.NewEncoder(w).Encode(user)
 } else{
	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(insertresult)
	json.NewEncoder(w).Encode(user)
 }


}