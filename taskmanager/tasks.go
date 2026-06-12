package main

import (
	"fmt"
	"gogit/taskmanager/handlers"
	"gogit/taskmanager/repository"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from go!")
}

func main() {

	http.HandleFunc("/hello", helloHandler)
	repo := repository.NewTaskRepository()
	taskHandler := handlers.NewTaskHandler(repo)
	http.HandleFunc("/tasks", taskHandler.GetTasks)

	fmt.Println("Server is running on http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}
