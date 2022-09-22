package main

import (
	"log"
	"net/http"

	"github.com/Golang_assingment/pkg/Handlers"
	"github.com/gorilla/mux"
)



func main() {
	router := mux.NewRouter()

	router.HandleFunc("/books", Handlers.GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", Handlers.GetBook).Methods(http.MethodGet)
	router.HandleFunc("/books", Handlers.AddBook).Methods(http.MethodPost)
	router.HandleFunc("/books/{id}", Handlers.UpdateBook).Methods(http.MethodPut)
	router.HandleFunc("/books/{id}",Handlers.Deletebook).Methods(http.MethodDelete)
	log.Println("Api is running")
	http.ListenAndServe(":4000", router)
}
