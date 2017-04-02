package dbclient

import (
	"github.com/sadhal/contacts-be/userservice/model"
	"gopkg.in/mgo.v2"
	"log"
	"fmt"
	"time"
	"gopkg.in/mgo.v2/bson"
	"os"
)

/*
var mgoSession   *mgo.Session

// Creates a new session if mgoSession is nil i.e there is no active mongo session.
//If there is an active mongo session it will return a Clone
func GetMongoSession() *mgo.Session {
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.Dial(mongo_conn_str)
		if err != nil {
			log.Fatal("Failed to start the Mongo session")
		}
	}
	return mgoSession.Clone()
}
*/

type IMongoClient interface {
	OpenMongoDb()
	QueryUser(userId string) (model.User, error)
	QueryUsers() ([]model.User, error)
}

// Real implementation
type MongoClient struct {
	mongoSession *mgo.Session
}

var MongoDBHosts = os.Getenv("MONGODB_SERVICE_HOST") + ":" + os.Getenv("MONGODB_SERVICE_PORT")

//const (
//	AuthDatabase = "sampledb"
//	AuthUserName = "sadhal"
//	AuthPassword = "sadhal"
//)
var AuthDatabase = os.Getenv("MONGODB_DATABASE")
var AuthUserName = os.Getenv("MONGODB_USER")
var AuthPassword = os.Getenv("MONGODB_PASSWORD")


func (mc *MongoClient) OpenMongoDb()  {
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{MongoDBHosts},
		Timeout:  60 * time.Second,
		Database: AuthDatabase,
		Username: AuthUserName,
		Password: AuthPassword,
	}

	var err error
	mc.mongoSession, err = mgo.DialWithInfo(mongoDBDialInfo)

	if err != nil {
		log.Fatal(err)
	}

}

func (mc *MongoClient) QueryUser(userId string) (model.User, error) {
	// Allocate an empty User instance we'll let json.Unmarshal populate for us in a bit.
	user := model.User{}

	//defer mc.mongoDB.Close()
	//mc.mongoDB.SetMode(mgo.Monotonic, true)

	c := mc.mongoSession.DB("sampledb").C("personer")


	fmt.Println("Quering Mongodb for user with id", userId)
	err := c.FindId(bson.M{"_id": bson.ObjectIdHex(userId)}).One(&user)
	//err := c.Find(bson.M{"firstName": "fname1"}).One(&user)
	if err != nil {
		log.Fatal("error occured:", err)
	}



	// If there were an error, return the error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (mc *MongoClient) QueryUsers() ([]model.User, error) {
	// Allocate an empty User instance we'll let json.Unmarshal populate for us in a bit.
	var all []model.User

	//defer mc.mongoDB.Close()
	//mc.mongoDB.SetMode(mgo.Monotonic, true)

	c := mc.mongoSession.DB("sampledb").C("personer")

	fmt.Println("Quering Mongodb for all users")
	errQ := c.Find(nil).All(&all)
	if errQ != nil {
		fmt.Println("RunQuery : ERROR")
		log.Printf("RunQuery : ERROR : %s\n", errQ)
		return nil, errQ
	}

	return all, nil
}