package handlers

import (
        // native packages
        "encoding/json"
        "net/http"
        // third party packages
        "github.com/gorilla/mux"
        mgo "gopkg.in/mgo.v2"
        "gopkg.in/mgo.v2/bson"
        // my local models
        "github.com/joscelynjames/rest-api/models"
)

// Userhandler represents the handler for operating on the Userresource
type UserHandler struct {
        session *mgo.Session
}

func NewUserHandler(session *mgo.Session) *UserHandler {
        return &UserHandler{session}
}

func (uh UserHandler) AllUsers(w http.ResponseWriter, r *http.Request) {
        var results []models.User

        if err := uh.session.DB("wheres-my-bird-db").C("users").Find(nil).All(&results); err != nil {
                w.WriteHeader(404)
                return
        }

        // write content-type, status code payload
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(200)
        json.NewEncoder(w).Encode(results)
}

func (uh UserHandler) OneUser(w http.ResponseWriter, r *http.Request) {
        // get id
        vars := mux.Vars(r)
        id := vars["id"]
        // verify id is ObjectId, otherwise return
        if !bson.IsObjectIdHex(id) {
                w.WriteHeader(404)
                return
        }
        // get id
        oid := bson.ObjectIdHex(id)
        u := models.User{}
        // fetch user
        if err := uh.session.DB("wheres-my-bird-db").C("users").FindId(oid).One(&u); err != nil {
                w.WriteHeader(404)
                return
        }
        // write content-type, status code, payload
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(200)
        json.NewEncoder(w).Encode(u)
}

func (uh UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
        // Stub an user to be populated from the body
        u := models.User{}
        // Populate the user data
        json.NewDecoder(r.Body).Decode(&u)
        // Add an Id
        u.ID = bson.NewObjectId()
        // write the user to mongodb
        uh.session.DB("wheres-my-bird-db").C("users").Insert(u)
        // Write content-type, statuscode, payload
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(201)
        json.NewEncoder(w).Encode(u)
}