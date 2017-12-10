package handlers

import (
        // native packages
        "encoding/json"
		"net/http"
		"time"
		"fmt"
		// "strconv"
        // third party packages
        // "github.com/gorilla/mux"
        mgo "gopkg.in/mgo.v2"
        "gopkg.in/mgo.v2/bson"
        // my local models
        "github.com/joscelynjames/rest-api/models"
)

type CoordHandler struct {
        session *mgo.Session
}

func NewCoordHandler(session *mgo.Session) *CoordHandler {
        return &CoordHandler{session}
}

func (ch CoordHandler) AllCoords(w http.ResponseWriter, r *http.Request) {
        var results []models.Coordinate

        if err := ch.session.DB("wheres-my-bird-db").C("coordinates").Find(nil).All(&results); err != nil {
                w.WriteHeader(404)
                return
        }

        // write content-type, status code payload
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(200)
        json.NewEncoder(w).Encode(results)
}

func (ch CoordHandler) CreateCoord(w http.ResponseWriter, r *http.Request) {
		// get the coordinate struct from models
		c := models.Coordinate{}	
		// json decode the request body and put it into the coordinate struct
		logthis := json.NewDecoder(r.Body).Decode(&c)
		fmt.Println(logthis)
		// generate new bson object id
		c.ID = bson.NewObjectId()
		c.Time = time.Now().Local()
		// c.X = strconv.ParseInt(c.X, 10, 64)
		// open session and insert the new coordinate 
		ch.session.DB("wheres-my-bird-db").C("coordinates").Insert(c)
		// set response header content
		w.Header().Set("Content-Type", "application/json")
		// status code
		w.WriteHeader(201)
		// json encode the response 
		json.NewEncoder(w).Encode(c)
}