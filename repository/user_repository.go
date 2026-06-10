package repository

import (
	"encoding/json"
	"errors"
	"os"

	"gogit/models"
)

type UserStore interface {
	CreateUser(name string, email string) models.User
	GetUserByID(id int) (models.User, error)
	GetAllUsers() []models.User
	Save() error
	Load() error
}

type UserRepository struct {
	users map[int]models.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: make(map[int]models.User),
	}
}

func (repo *UserRepository) CreateUser(
	name string,
	email string,
) models.User {
	id := len(repo.users) + 1

	user := models.User{
		ID:    id,
		Name:  name,
		Email: email,
	}

	repo.AddUser(user)

	return user
}

func (repo *UserRepository) AddUser(user models.User) {
	repo.users[user.ID] = user
}

func (repo *UserRepository) GetAllUsers() []models.User {
	var users []models.User
	for _, user := range repo.users {
		users = append(users, user)
	}
	return users
}

func (repo *UserRepository) GetUserByID(id int) (models.User, error) {
	user, exists := repo.users[id]

	if !exists {
		return models.User{}, errors.New("user not found")
	}

	return user, nil
}

func (repo *UserRepository) Save() error {
	filename := "users.json"
	var users []models.User

	for _, user := range repo.users {
		users = append(users, user)
	}

	data, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

func (repo *UserRepository) Load() error {
	filename := "users.json"
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	var users []models.User

	err = json.Unmarshal(data, &users)
	if err != nil {
		return err
	}

	for _, user := range users {
		repo.users[user.ID] = user
	}

	return nil
}
