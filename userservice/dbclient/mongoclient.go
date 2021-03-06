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

type IMongoClient interface {
	OpenMongoDb()
	QueryUser(userId string) (model.User, error)
	QueryUsers() ([]model.User, error)
	CreateUser(user *model.User) (model.User, error)
}

// Real implementation
type MongoClient struct {
	mongoSession *mgo.Session
}

var MongoDBHosts = os.Getenv("MONGODB_SERVICE_HOST") + ":" + os.Getenv("MONGODB_SERVICE_PORT")

const (
	AuthDatabase = "sampledb"
	AuthUserName = "sadhal"
	AuthPassword = "sadhal"
)
//var AuthDatabase = os.Getenv("MONGODB_DATABASE")
//var AuthUserName = os.Getenv("MONGODB_USER")
//var AuthPassword = os.Getenv("MONGODB_PASSWORD")


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

	session := mc.mongoSession.Copy()
	defer session.Close()
	//defer mc.mongoDB.Close()
	session.SetMode(mgo.Monotonic, true)

	//c := mc.mongoSession.DB("sampledb").C("personer")
	c := session.DB("sampledb").C("personer")


	fmt.Println("Quering Mongodb for user with id", userId)
	//err := c.FindId(bson.M{"_id": bson.ObjectIdHex(userId)}).One(&user)
	//err := c.Find(bson.M{"firstName": "firstName_189"}).One(&user)
	//err := c.FindId(userId).One(&user)
	//err := c.Find(bson.M{"_id": userId}).One(&user)
	err := c.FindId(bson.ObjectIdHex(userId)).One(&user)
	if err != nil {
		fmt.Println("error occured:", err)
		qerr, other := err.(*mgo.QueryError)
		fmt.Println("qerr: ", qerr)
		fmt.Println("other: ", other)
		lerr, other2 := err.(*mgo.LastError)
		fmt.Println("lerr: ", lerr)
		fmt.Println("other2: ", other2)
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

	session := mc.mongoSession.Copy()
	defer session.Close()
	//defer mc.mongoDB.Close()
	session.SetMode(mgo.Monotonic, true)

	//c := mc.mongoSession.DB("sampledb").C("personer")
	c := session.DB("sampledb").C("personer")

	fmt.Println("Quering Mongodb for all users")
	errQ := c.Find(nil).All(&all)
	if errQ != nil {
		fmt.Println("RunQuery : ERROR")
		log.Printf("RunQuery : ERROR : %s\n", errQ)
		return nil, errQ
	}

	return all, nil
}

func (mc *MongoClient) CreateUser(user *model.User) (model.User, error) {
	session := mc.mongoSession.Copy()
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	//c := mc.mongoSession.DB("sampledb").C("personer")
	c := session.DB("sampledb").C("personer")
	fmt.Println("Creating user in Mongodb: ", user)

	// Add an Id
	user.Id = bson.NewObjectId()

	err := c.Insert(user)
	// If there were an error, return the error
	if err != nil {
		log.Fatal("error occured:", err.Error())
		return model.User{}, err
	}
	return *user, nil
}