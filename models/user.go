package models

import "gopkg.in/mgo.v2/bson"

type User struct {
        ID       bson.ObjectId `json:"id" bson:"_id,omitempty"`
        Name     string        `json:"name" bson:"name"`
        Email    string        `json:"email" bson:"email"`
        Password string        `json:"password" bson:"password"`
        Admin    bool          `json:"admin" bson:"admin"`
}