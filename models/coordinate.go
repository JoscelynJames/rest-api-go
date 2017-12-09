package models

import (
        "gopkg.in/mgo.v2/bson"
        "time"
)

// coordinates
type Coordinate struct {
        ID   bson.ObjectId `json:"id"`
        X    int           `json:"x"`
        Y    int           `json:"y"`
        Time time.Time     `json:"time"`
}