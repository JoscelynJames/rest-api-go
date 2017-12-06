package main

import (
    "fmt"
    "log"
		"net/http"
		"encoding/json"
		"time"

		"github.com/gorilla/mux"
		// "github.com/joscelynjames/rest-api/model"
)

type User struct {
    Name      string
    Admin     bool
    Created   time.Time
}

type Users []User

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/users", UserIndex)
	router.HandleFunc("/users/{userId}", UserShow)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}

func UserIndex(w http.ResponseWriter, r *http.Request) {
	users := Users{
		User{Name: "Write presentation"},
		User{Name: "Host meetup"},
	}

	json.NewEncoder(w).Encode(users)
}

func UserShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    userId := vars["userId"]
    fmt.Fprintln(w, "User show:", userId)
}