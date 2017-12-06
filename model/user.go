package model

import "time"

type User struct {
    Name      string
    Admin     bool
    Created   time.Time
}

type Users []User