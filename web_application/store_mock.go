package main

import (
	"github.com/stretchr/testify/mock"
)

// The mock store contains additional methods for inspection
type MockStore struct {
	mock.Mock
}

func (m *MockStore) CreateBird(bird *Bird) error {
	/*
		when the method is called, `m.called` records the call, and also returns the result
		that we pass to it (which you will see in the handler tests)
	*/
	rets := m.Called(bird)
	return rets.Error(0)
}

func (m *MockStore) GetBirds() ([]*Bird, error) {
	rets := m.Called()
	/*
		Since the rets.Get() is a generic method, that returns whaever we pass to it, we need to type cast it to the type we expect, which in this case is []*Bird
	*/
	return rets.Get(0).([]*Bird), rets.Error(1)
}

func InitMockStore() *MockStore {
	/*
		Like the InitStore function we defined earlier, this function
		also initializes the store variable, but this time, it assigns
		a new MockStore instance to it, instead of an actual store
	*/
	s := new(MockStore)
	store = s
	return s
}
