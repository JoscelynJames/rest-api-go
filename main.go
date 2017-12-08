package main

import (
    "log"
		"net/http"
		"encoding/json"
		"strconv"

		"github.com/gorilla/mux"
		//future use 
		// "github.com/joscelynjames/rest-api/model"
)
// users
type User struct {
		ID        int    `json:"id"` 
		Name      string `json:"name,"`
		Email     string `json:"email,"`
		Password  string `json:"password,"`
		Admin     bool   `json:"admin,"`
}
var users []User

// coordinates 
type Coordinate struct {
	ID    int    `json:"id"`
	X 		int 	 `json:"x"`
	Y 		int 	 `json:"y"`
	Time  string `json:"time"`
}
var coordinates []Coordinate

func main() {
	router := mux.NewRouter().StrictSlash(true)

	users = append(users, User{ID: 1, Name: "Joscelyn Jmames", Email: "josce.james7@gmail.com", Password: "missy7", Admin: true })
	users = append(users, User{ID: 2, Name: "Bradford Lamson-Scribner", Email: "bradford@gmail.com", Password: "miles", Admin: true })
	users = append(users, User{ID: 3, Name: "Cindy Ann James", Email: "CindyAnnJames@gmail.com", Password: "catsrule", Admin: false })
	users = append(users, User{ID: 4, Name: "Jacqueline Jmames", Email: "jac-attac@gmail.com", Password: "charles123", Admin: false })
	users = append(users, User{ID: 5, Name: "Joel Jmames", Email: "jjames@gmail.com", Password: "niko80", Admin: false })

	coordinates = append(coordinates, Coordinate{ID: 1, X: 5, Y: 8, Time: "December 23 20017" })
	coordinates = append(coordinates, Coordinate{ID: 2, X: 5, Y: 2, Time: "December 23 20017" })
	coordinates = append(coordinates, Coordinate{ID: 3, X: 3, Y: 4, Time: "December 23 20017" })
	coordinates = append(coordinates, Coordinate{ID: 4, X: 2, Y: 4, Time: "December 23 20017" })
	coordinates = append(coordinates, Coordinate{ID: 5, X: 2, Y: 6, Time: "December 23 20017" })
	coordinates = append(coordinates, Coordinate{ID: 6, X: 3, Y: 5, Time: "December 23 20017" })
	coordinates = append(coordinates, Coordinate{ID: 7, X: 4, Y: 6, Time: "December 23 20017" })
	coordinates = append(coordinates, Coordinate{ID: 8, X: 4, Y: 7, Time: "December 23 20017" })
	coordinates = append(coordinates, Coordinate{ID: 9, X: 5, Y: 5, Time: "December 23 20017" })
	coordinates = append(coordinates, Coordinate{ID: 10, X: 6, Y: 7, Time: "December 23 20017" })
	coordinates = append(coordinates, Coordinate{ID: 11, X: 6, Y: 8, Time: "December 23 20017" })
	coordinates = append(coordinates, Coordinate{ID: 12, X: 7, Y: 7, Time: "December 23 20017" })
	coordinates = append(coordinates, Coordinate{ID: 13, X: 7, Y: 6, Time: "December 23 20017" })
	coordinates = append(coordinates, Coordinate{ID: 14, X: 10, Y: 5, Time: "December 23 20017" })

	router.HandleFunc("/users", allUsers)
	router.HandleFunc("/users/{userId}", oneUser)
	router.HandleFunc("/coords", allCoords)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func allUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

func oneUser(w http.ResponseWriter, r *http.Request) {
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

func allCoords(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(coordinates)
}


