package main

import (
    // "fmt"
    "log"
		"net/http"
		"encoding/json"
		"strconv"

		"github.com/gorilla/mux"
		//future use 
		// "github.com/joscelynjames/rest-api/model"
)

type User struct {
		ID        int
		Name      string
		Email     string
		Password  string
		Admin     bool
}

var users []User

func main() {

	router := mux.NewRouter().StrictSlash(true)

	users = append(users, User{ID: 1, Name: "Joscelyn Jmames", Email: "josce.james7@gmail.com", Password: "missy7", Admin: true })
	users = append(users, User{ID: 2, Name: "Bradford Lamson-Scribner", Email: "bradford@gmail.com", Password: "miles", Admin: true })
	users = append(users, User{ID: 3, Name: "Cindy Ann James", Email: "CindyAnnJames@gmail.com", Password: "catsrule", Admin: false })
	users = append(users, User{ID: 4, Name: "Jacqueline Jmames", Email: "jac-attac@gmail.com", Password: "charles123", Admin: false })
	users = append(users, User{ID: 5, Name: "Joel Jmames", Email: "jjames@gmail.com", Password: "niko80", Admin: false })

	router.HandleFunc("/users", UserIndex)
	router.HandleFunc("/users/{userId}", UserShow)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func UserIndex(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

func UserShow(w http.ResponseWriter, r *http.Request) {
	// get params from the request
		vars := mux.Vars(r)
	// set the param user id to var 
		intId := vars["userId"]
	// convert the string to a number 
		userId, err := strconv.Atoi(intId)
		if err != nil {
			panic(err)
		}
	// loop over users
    for _, user := range users {
			// find the user that id == the param
        if user.ID == userId {
            json.NewEncoder(w).Encode(user)
            return
        }
    }
    return 
}