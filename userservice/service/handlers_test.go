package service

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"net/http/httptest"
	"github.com/sadhal/contacts-be/userservice/dbclient"
	"github.com/sadhal/contacts-be/userservice/model"
	"fmt"
	"encoding/json"
	"time"
	"gopkg.in/mgo.v2/bson"
	"bytes"
)

func TestGetUsertWrongPath(t *testing.T) {

	Convey("Given a HTTP request for /invalid/123", t, func() {
		req := httptest.NewRequest("GET", "/invalid/123", nil)
		resp := httptest.NewRecorder()

		Convey("When the request is handled by the Router", func() {
			NewRouter().ServeHTTP(resp, req)

			Convey("Then the response should be a 404", func() {
				So(resp.Code, ShouldEqual, 404)
			})
		})
	})
}

func TestGetUser(t *testing.T) {
	// Create a mock instance that implements the IMongoClient interface
	mockRepo := &dbclient.MockMongoClient{}

	// Declare two mock behaviours. For "123" as input, return a proper Account struct and nil as error.
	// For "456" as input, return an empty User object and a real error.
	mockRepo.On("QueryUser", "123").Return(model.User{
		"fname123",
		"lname123",
		"e123@a.se",
		"tweeeeet123",
		time.Now(),
		bson.ObjectIdHex("56bdd27ecfa93bfe3d35047d"),
	}, nil)
	mockRepo.On("QueryUser", "456").Return(model.User{}, fmt.Errorf("Some error"))

	// Finally, assign mockRepo to the DBClient field (it's in _handlers.go_, e.g. in the same package)
	MongoDbClient = mockRepo
	Convey("Given a HTTP request for /personer/123", t, func() {
		req := httptest.NewRequest("GET", "/personer/123", nil)
		resp := httptest.NewRecorder()

		Convey("When the request is handled by the Router", func() {
			NewRouter().ServeHTTP(resp, req)

			Convey("Then the response should be a 200", func() {
				So(resp.Code, ShouldEqual, 200)

				user := model.User{}
				json.Unmarshal(resp.Body.Bytes(), &user)
				So(user.Id, ShouldEqual, bson.ObjectIdHex("56bdd27ecfa93bfe3d35047d"))
				So(user.FirstName, ShouldEqual, "fname123")
			})
		})
	})

	Convey("Given a HTTP request for /personer/456", t, func() {
		req := httptest.NewRequest("GET", "/personer/456", nil)
		resp := httptest.NewRecorder()

		Convey("When the request is handled by the Router", func() {
			NewRouter().ServeHTTP(resp, req)

			Convey("Then the response should be a 404", func() {
				So(resp.Code, ShouldEqual, 404)
			})
		})
	})
}

//func TestHello(t *testing.T) {
//	Convey("Given a call to function Hello", t, func() {
//		Convey("When the function is called", func() {
//			Convey("Then the response should be 'Hello, World!'", func() {
//				const correct = "Hello, world!"
//				So(Hello(), ShouldEqual, correct)
//			})
//		})
//	})
//}

func TestCreateUserSadPath(t *testing.T) {
	// Create a mock instance that implements the IMongoClient interface
	mockRepo := &dbclient.MockMongoClient{}

	timeNow := time.Now()

	nokUser := &model.User{ FirstName: "notOkUser", CreatedOn: timeNow}

	// Sad path
	mockRepo.On("CreateUser", nokUser).Return(model.User{}, fmt.Errorf("Some error"))

	// Finally, assign mockRepo to the DBClient field (it's in _handlers.go_, e.g. in the same package)
	MongoDbClient = mockRepo

	Convey("Given a unsuccessful HTTP request for /personer", t, func() {
		nokUserBytes, _ := json.Marshal(nokUser)
		reader := bytes.NewReader(nokUserBytes)
		//reader := strings.NewReader(nokUserString)
		req := httptest.NewRequest("POST", "/personer", reader)
		resp := httptest.NewRecorder()

		Convey("When the request is handled by the Router", func() {
			NewRouter().ServeHTTP(resp, req)

			Convey("Then the response should be a 500", func() {
				So(resp.Code, ShouldEqual, 500)
			})
		})
	})
}

func TestCreateUserHappyPath(t *testing.T) {
	// Create a mock instance that implements the IMongoClient interface
	mockRepo := &dbclient.MockMongoClient{}

	timeNow := time.Now()
	objectId := bson.ObjectIdHex("58e2211e33af735cba906d34")

	// Finally, assign mockRepo to the DBClient field (it's in _handlers.go_, e.g. in the same package)
	MongoDbClient = mockRepo

	Convey("Given a successful HTTP request for /personer", t, func() {
		// Test data
		okUser := &model.User{
			FirstName: 	"fname123",
			LastName:	"lname123",
			Email: 		"e123@a.se",
			TwitterHandle: 	"tweeeeet123",
			CreatedOn: 	timeNow,
			//Id: 		objectId,
		}

		//// Happy path mock
		mockRepo.On("CreateUser", okUser).Return(model.User{ FirstName: "fname123",
			LastName: okUser.LastName,
			CreatedOn: okUser.CreatedOn,
			Id: objectId,
		}, nil)

		okUserBytes, _ := json.Marshal(okUser)
		reader := bytes.NewReader(okUserBytes)
		req := httptest.NewRequest("POST", "/personer", reader)
		resp := httptest.NewRecorder()

		Convey("When the request is handled by the Router", func() {
			NewRouter().ServeHTTP(resp, req)

			Convey("Then the response should be a 201", func() {
				So(resp.Code, ShouldEqual, 201)

				user := model.User{}
				json.Unmarshal(resp.Body.Bytes(), &user)
				So(user.FirstName, ShouldEqual, "fname123")
				So(user.Id, ShouldEqual, objectId)
			})
		})
	})

	Convey("Given a successful HTTP request 2 for /personer", t, func() {
		//// Happy path with another syntax
		mockRepo.On("CreateUser", &model.User{FirstName: "yo", Email: "e@a.se"}).Return(model.User{
			"yo",
			"oy",
			"e@a.se",
			"tweeeeet123",
			time.Now(),
			bson.ObjectIdHex("56bdd27ecfa93bfe3d35047d"),
		}, nil)

		// Test data
		payload := map[string]interface{}{
			"firstName":	"yo",
			"email":   	"e@a.se",
		}
		
		jsonPayload, _ := json.Marshal(payload)
		reader := bytes.NewReader(jsonPayload)
		reader.Len()
		req := httptest.NewRequest("POST", "/personer", reader)
		resp := httptest.NewRecorder()

		Convey("When the request is handled by the Router", func() {
			NewRouter().ServeHTTP(resp, req)

			Convey("Then the response should be a 201", func() {
				So(resp.Code, ShouldEqual, 201)

				user := model.User{}
				json.Unmarshal(resp.Body.Bytes(), &user)
				So(user.FirstName, ShouldEqual, "yo")
				So(user.Id, ShouldEqual, bson.ObjectIdHex("56bdd27ecfa93bfe3d35047d"))
			})
		})
	})
}