package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

// coordinates
type Coordinate struct {
	ID   bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Time time.Time     `json:"time" bson:"time"`
	X    int           `json:"x" bson:"x"`
	Y    int           `json:"y" bson:"y"`
}
