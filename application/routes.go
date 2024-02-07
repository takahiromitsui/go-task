package application

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/takahiromitsui/go-task-manager/controller"
	"github.com/takahiromitsui/go-task-manager/repository/task"
)

func(app *App) loadRoutes() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	app.loadTaskRoutes(router)
	app.router = router
}

func (app *App) loadTaskRoutes(router chi.Router) {
	taskController := &controller.Task{
		Repository: &task.PostgresRepository{
			DB: app.db,
		},
	}
	router.Post("/tasks", taskController.Create)
	
}