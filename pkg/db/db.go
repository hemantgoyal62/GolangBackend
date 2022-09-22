package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
func Client()(*mongo.Collection){
	//establishing connection with mysql cluster
	clientOptions := options.Client().ApplyURI("mongodb+srv://hemant:hemant@cluster0.1ccb4.mongodb.net/?retryWrites=true&w=majority")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	
	if err!= nil{
		log.Fatal(err)
	}else{
		log.Print("Connection is successful")
	}
	
	
	
	err = client.Ping(context.TODO(),nil)
	if err!= nil{
		log.Print("fatal error due to ping")
		log.Fatal(err)
	}else{
		log.Print("Ping is successful")
	}
	// defer client.Disconnect(context.TODO())
	collection :=client.Database("Golang_assignment").Collection("Users")
	return collection
}