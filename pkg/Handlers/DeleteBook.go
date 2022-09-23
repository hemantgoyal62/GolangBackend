package Handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	db "github.com/Golang_assingment/pkg/db"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

func Deletebook(w http.ResponseWriter,r *http.Request) {
	vars:= mux.Vars(r)
	id,_:=strconv.Atoi(vars["id"])
	collection := db.Client()
	deleteResult,err:= collection.DeleteOne(context.TODO(),bson.D{{Key: "id",Value: id}})
	if err!= nil{
		log.Print(err)
	}else{
		w.Header().Add("Content-type","application/json")
		json.NewEncoder(w).Encode(deleteResult)
	}

}