package handlers

import (
        // native packages
        "net/http"

        // third party packages
        mgo "gopkg.in/mgo.v2"
)

type CoordHandler struct {
        session *mgo.Session
}

func NewCoordHandler(session *mgo.Session) *CoordHandler {
        return &CoordHandler{session}
}

func (ch CoordHandler) AllCoords(w http.ResponseWriter, r *http.Request) {
        // json.NewEncoder(w).Encode(coordinates)
}

func (ch CoordHandler) CreateCoord(w http.ResponseWriter, r *http.Request) {
        // json.NewEncoder(w).Encode(coordinates)
}