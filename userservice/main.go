package main

import (
	"fmt"
	"github.com/sadhal/contacts-be/userservice/service"
	"github.com/sadhal/contacts-be/userservice/dbclient"
)

var appName = "userservice"

func main() {
	fmt.Printf("Starting %v\n", appName)
	fmt.Printf("Connecting to Mongodb")

	initializeMongoClient()
	service.StartWebServer("6767")
}
func initializeMongoClient() {
	service.MongoDbClient = &dbclient.MongoClient{}
	service.MongoDbClient.OpenMongoDb()
}
