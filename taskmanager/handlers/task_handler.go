package handlers

import (
	"encoding/json"
	"gogit/taskmanager/repository"
	"net/http"
)

type TaskHandler struct {
	repo *repository.TaskRepository
}

func NewTaskHandler(repo *repository.TaskRepository) *TaskHandler {
	return &TaskHandler{
		repo: repo,
	}
}

func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks := h.repo.GetAllTasks()
	json.NewEncoder(w).Encode(tasks)
}
