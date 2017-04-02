package service

import (
	"github.com/sadhal/contacts-be/userservice/dbclient"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"strconv"
	"log"
)

//var DBClient dbclient.IBoltClient
var MongoDbClient dbclient.IMongoClient

func GetUser(w http.ResponseWriter, r *http.Request)  {
	// Read the 'userId' path parameter from the mux map
	var userId = mux.Vars(r)["userId"]

	// Read the user struct MongoDB
	user, err := MongoDbClient.QueryUser(userId)

	// If err, return a 404
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	log.Println("found user with id: ", userId)

	// If found, marshal into JSON, write headers and content
	data, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func GetUsers(w http.ResponseWriter, r *http.Request)  {

	// Read the user struct MongoDB
	users, err := MongoDbClient.QueryUsers()

	// If err, return a 404
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	log.Println("found users: ", len(users))

	// If found, marshal into JSON, write headers and content
	data, _ := json.Marshal(users)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}