package models

// coordinates 
type Coordinate struct {
	ID    int    `json:"id"`
	X 		int 	 `json:"x"`
	Y 		int 	 `json:"y"`
	Time  string `json:"time"`
}

var coordinates []Coordinate
