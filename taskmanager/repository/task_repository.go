package repository

import (
	"fmt"
	"gogit/taskmanager/models"
)

type TaskRepository struct {
	tasks map[int]models.Task
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{
		tasks: make(map[int]models.Task),
	}
}

func (repo *TaskRepository) CreateTask(title string) models.Task {
	id := len(repo.tasks) + 1
	task := models.Task{
		ID:        id,
		Title:     title,
		Completed: false,
	}

	repo.tasks[id] = task
	return task
}

func (repo *TaskRepository) GetAllTasks() []models.Task {
	tasks := []models.Task{}
	for _, task := range repo.tasks {
		tasks = append(tasks, task)
	}
	fmt.Println(tasks)
	return tasks
}
