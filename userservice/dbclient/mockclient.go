package dbclient

import (
	"github.com/stretchr/testify/mock"
	"github.com/sadhal/contacts-be/userservice/model"
)

// MockBoltClient is a mock implementation of a datastore client for testing purposes.
// Instead of the bolt.DB pointer, we're just putting a generic mock object from
// strechr/testify
type MockMongoClient struct {
	mock.Mock
}

// From here, we'll declare three functions that makes our MockBoltClient fulfill the interface IBoltClient that we declared in part 3.
func (m *MockMongoClient) QueryUser(userId string) (model.User, error) {
	args := m.Mock.Called(userId)
	return args.Get(0).(model.User), args.Error(1)
}

func (m *MockMongoClient) QueryUsers() ([]model.User, error) {
	args := m.Mock.Called()
	return args.Get(0).([]model.User), args.Error(1)
}

func (m *MockMongoClient) OpenMongoDb() {
	// Does nothing
}

func (m *MockMongoClient) CreateUser(user *model.User) (model.User, error) {
	args := m.Mock.Called(user)
	return args.Get(0).(model.User), args.Error(1)
}

