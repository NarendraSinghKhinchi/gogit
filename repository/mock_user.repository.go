package repository

import "gogit/models"

type MockUserRepository struct {
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{}
}

func (repo *MockUserRepository) CreateUser(name string, email string) models.User {
	return models.User{
		ID:    1,
		Name:  name,
		Email: email,
	}
}

func (repo *MockUserRepository) GetUserByID(id int) (models.User, error) {
	return models.User{
		ID:    id,
		Name:  "Mock User",
		Email: "mock.user@gmail.com",
	}, nil
}

func (repo *MockUserRepository) GetAllUsers() []models.User {
	return []models.User{}
}

func (repo *MockUserRepository) Save() error {
	return nil
}

func (repo *MockUserRepository) Load() error {
	return nil
}
