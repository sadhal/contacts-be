package model

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	FirstName 	string		`json:"firstName"`
	LastName 	string		`json:"lastName"`
	Email		string		`json:"email"`
	TwitterHandle	string		`json:"twitterHandle"`
	CreatedOn 	time.Time	`json:"createdOn"`
	Id		bson.ObjectId 	`bson:"_id,omitempty"`
}
