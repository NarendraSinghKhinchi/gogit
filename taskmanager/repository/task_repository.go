package repository

import (
	"fmt"
	"gogit/taskmanager/models"
)

type TaskStore interface {
	GetTaskByID(id int) (models.Task, bool)
	GetAllTasks() []models.Task
	CreateTask(title string) models.Task
	DeleteTaskByID(id int) bool
}
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

func (repo *TaskRepository) GetTaskByID(id int) (models.Task, bool) {
	task, exists := repo.tasks[id]
	return task, exists
}

func (repo *TaskRepository) DeleteTaskByID(id int) bool {
	if _, exists := repo.tasks[id]; exists {
		delete(repo.tasks, id)
		return true
	}
	return false
}
