package main

import (
	"fmt"
	"gogit/repository"
)

func main() {
	var store repository.UserStore
	store = repository.NewMockUserRepository()
	err := store.Load()
	if err != nil {
		fmt.Println("Error loading from file:", err)
	}

	fmt.Println("All Users:")
	for _, user := range store.GetAllUsers() {
		fmt.Printf(
			"User: %s, ID: %d, Email: %s\n",
			user.Name,
			user.ID,
			user.Email,
		)
	}

	user, err := store.GetUserByID(1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Found:", user.Name)

	err = store.Save()
	if err != nil {
		fmt.Println("Error saving file:", err)
	}
}
