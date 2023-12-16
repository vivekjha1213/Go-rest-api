package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const PREFIX = "/users"

func initializerRouter() {
	r := mux.NewRouter()

	r.HandleFunc(PREFIX, GetUsers).Methods("GET")
	r.HandleFunc(PREFIX+"/{id}", GetUser).Methods("GET")
	r.HandleFunc(PREFIX, CreateUser).Methods("POST")
	r.HandleFunc(PREFIX+"/{id}", UpdateUser).Methods("PUT")
	r.HandleFunc(PREFIX+"/{id}", DeleteUser).Methods("DELETE")
	


	log.Fatal(http.ListenAndServe(":9000",r))

}

func main(){
    InitialMigration()
	initializerRouter()
	

}
