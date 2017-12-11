package main

import (
	"os"
	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
	"log"
	"net/http"
	"github.com/joscelynjames/rest-api/handlers"
)

var db *mgo.Session

func getSession() *mgo.Session {
	// Connect to our local mongo
	db, err := mgo.Dial(os.Getenv("PROD_MONGODB"))
	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	return db
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	uh := handlers.NewUserHandler(getSession())
	ch := handlers.NewCoordHandler(getSession())

	router.HandleFunc("/users", uh.AllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", uh.OneUser).Methods("GET")
	router.HandleFunc("/user", uh.CreateUser).Methods("POST")

	router.HandleFunc("/coords", ch.AllCoords).Methods("GET")
	router.HandleFunc("/coord", ch.CreateCoord).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
