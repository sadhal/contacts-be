package model

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	FirstName 	string		`json:"firstName" bson:"firstName"`
	LastName 	string		`json:"lastName" bson:"lastName"`
	Email		string		`json:"email"`
	TwitterHandle	string		`json:"twitterHandle" bson:"twitterHandle"`
	CreatedOn 	time.Time	`json:"createdOn,Date" bson:"createdDate"`
	Id		bson.ObjectId 	`json:"id" bson:"_id,omitempty"`
}
