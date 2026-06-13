package handlers

import (
	"encoding/json"
	"gogit/taskmanager/repository"
	"net/http"
	"strconv"
)

type TaskHandler struct {
	repo repository.TaskStore
}

func NewTaskHandler(repo repository.TaskStore) *TaskHandler {
	return &TaskHandler{
		repo: repo,
	}
}

func (h *TaskHandler) getTasks(w http.ResponseWriter, r *http.Request) {
	tasks := h.repo.GetAllTasks()
	json.NewEncoder(w).Encode(tasks)
}

func (h *TaskHandler) getTask(w http.ResponseWriter, r *http.Request) {

	taskID := r.URL.Query().Get("id")

	id, err := strconv.Atoi(taskID)

	if err != nil {
		http.Error(w, "Invalid task id", http.StatusBadRequest)
		return
	}

	task, exists := h.repo.GetTaskByID(id)

	if !exists {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(task)
}

type CreateTaskRequest struct {
	Title string `json:"title"`
}

func (h *TaskHandler) createTask(w http.ResponseWriter, r *http.Request) {
	var req CreateTaskRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	task := h.repo.CreateTask(req.Title)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) deleteTask(w http.ResponseWriter, r *http.Request) {
	taskID := r.URL.Query().Get("id")
	id, err := strconv.Atoi(taskID)
	if err != nil {
		http.Error(w, "Invalid task id", http.StatusBadRequest)
		return
	}

	if h.repo.DeleteTaskByID(id) {
		w.WriteHeader(http.StatusNoContent)
	} else {
		http.Error(w, "Task not found", http.StatusNotFound)
	}
}

func (h *TaskHandler) Tasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		taskID := r.URL.Query().Get("id")
		if taskID != "" {
			h.getTask(w, r)
		} else {
			h.getTasks(w, r)
		}

	case http.MethodPost:
		h.createTask(w, r)

	case http.MethodDelete:
		h.deleteTask(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
