package main

import (
    "fmt"
    "log"
		"net/http"
		// "time"
		mgo "gopkg.in/mgo.v2"
		"github.com/gorilla/mux"
		//future use 
		"github.com/joscelynjames/rest-api/handlers"
)

var db *mgo.Session

func getSession() *mgo.Session {  
    // Connect to our local mongo
    db, err := mgo.Dial("mongodb://localhost")

    // Check if connection error, is mongo running?
    if err != nil {
        panic(err)
		}
		fmt.Print("Connected to mongo")
    return db
}

func main() {
	getSession()
	router := mux.NewRouter().StrictSlash(true)
	uh := handlers.NewUserHandler(getSession())

	router.HandleFunc("/users", uh.AllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", uh.OneUser).Methods("GET")
	router.HandleFunc("/coords", uh.AllCoords).Methods("GET")
	router.HandleFunc("/user", uh.CreateUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}

