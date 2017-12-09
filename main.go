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

	// users = append(users, User{ID: 1, Name: "Joscelyn Jmames", Email: "josce.james7@gmail.com", Password: "missy7", Admin: true })
	// users = append(users, User{ID: 2, Name: "Bradford Lamson-Scribner", Email: "bradford@gmail.com", Password: "miles", Admin: true })
	// users = append(users, User{ID: 3, Name: "Cindy Ann James", Email: "CindyAnnJames@gmail.com", Password: "catsrule", Admin: false })
	// users = append(users, User{ID: 4, Name: "Jacqueline Jmames", Email: "jac-attac@gmail.com", Password: "charles123", Admin: false })
	// users = append(users, User{ID: 5, Name: "Joel Jmames", Email: "jjames@gmail.com", Password: "niko80", Admin: false })

	// coordinates = append(coordinates, Coordinate{ID: 1, X: 5, Y: 8, Time: "December 23 20017" })
	// coordinates = append(coordinates, Coordinate{ID: 2, X: 5, Y: 2, Time: "December 23 20017" })
	// coordinates = append(coordinates, Coordinate{ID: 3, X: 3, Y: 4, Time: "December 23 20017" })
	// coordinates = append(coordinates, Coordinate{ID: 4, X: 2, Y: 4, Time: "December 23 20017" })
	// coordinates = append(coordinates, Coordinate{ID: 5, X: 2, Y: 6, Time: "December 23 20017" })
	// coordinates = append(coordinates, Coordinate{ID: 6, X: 3, Y: 5, Time: "December 23 20017" })
	// coordinates = append(coordinates, Coordinate{ID: 7, X: 4, Y: 6, Time: "December 23 20017" })
	// coordinates = append(coordinates, Coordinate{ID: 8, X: 4, Y: 7, Time: "December 23 20017" })
	// coordinates = append(coordinates, Coordinate{ID: 9, X: 5, Y: 5, Time: "December 23 20017" })
	// coordinates = append(coordinates, Coordinate{ID: 10, X: 6, Y: 7, Time: "December 23 20017" })
	// coordinates = append(coordinates, Coordinate{ID: 11, X: 6, Y: 8, Time: "December 23 20017" })
	// coordinates = append(coordinates, Coordinate{ID: 12, X: 7, Y: 7, Time: "December 23 20017" })
	// coordinates = append(coordinates, Coordinate{ID: 13, X: 7, Y: 6, Time: "December 23 20017" })
	// coordinates = append(coordinates, Coordinate{ID: 14, X: 10, Y: 5, Time: "December 23 20017" })
	uh := handlers.NewUserHandler(getSession())

	router.HandleFunc("/users", uh.AllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", uh.OneUser).Methods("GET")
	router.HandleFunc("/coords", uh.AllCoords).Methods("GET")
	router.HandleFunc("/user", uh.CreateUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}


