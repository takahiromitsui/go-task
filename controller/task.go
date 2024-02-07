package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/takahiromitsui/go-task-manager/model"
	"github.com/takahiromitsui/go-task-manager/repository/task"
)

type Task struct {
	Repository *task.PostgresRepository
}

func (t *Task) Create(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		fmt.Print(r.Body)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	task := model.Task{
		Title:       body.Title,
		Description: body.Description,
		IsDone:      false,
	}
	if err := t.Repository.Insert(r.Context(), task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}