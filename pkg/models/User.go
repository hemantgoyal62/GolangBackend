package models

import (
	"time"
)


type User struct {
	Id        int    `json:"id" bson:"id"`
	Name      string `json:"name" bson:"name"`
	DOB       string `json:"dob" bson:"dob"`
	Address   string `json:"address" bson:"address"`
	Desc      string `json:"desc" bson:"desc"`
	CreatedAt time.Time `json:"createdAt omitempty" bson:"createdAt omitempty"`
}