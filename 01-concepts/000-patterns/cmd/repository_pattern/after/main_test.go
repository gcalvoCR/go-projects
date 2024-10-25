package main

import "testing"

type MockUserRepository struct{}

func (m *MockUserRepository) GetUserByID(id int) (User, error) {
	return User{ID: id, Name: "Test User"}, nil
}

func TestMain(t *testing.T) {
	// Use the mock repository for testing
	mockRepo := &MockUserRepository{}
	GetUserProfile(mockRepo, 1)
}
