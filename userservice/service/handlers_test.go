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