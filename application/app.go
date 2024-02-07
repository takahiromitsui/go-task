package application

import (
	"context"
	"fmt"
	"net/http"

	// "os"
	"time"

	_ "github.com/lib/pq"
	"github.com/takahiromitsui/go-task-manager/util"
	"gorm.io/gorm"
)


type App struct {
	router http.Handler
	db 	 *gorm.DB
}

func init() {
	util.LoadEnv()
	util.ConnectToDB()
}

func NewApp() *App {
	app:= &App{
		db: util.DB,
	}
	app.loadRoutes()
	return app
}

func (app *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":8080",
		Handler: app.router,
	}

	fmt.Println("Server is running on port 8080")
	ch := make(chan error, 1)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("failed to start server: %v", err)
		}
		close(ch)
	}()

	select {
		case err := <-ch:
			return err
		case <-ctx.Done():
			timeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			if err := server.Shutdown(timeout); err != nil {
				return fmt.Errorf("failed to shutdown server: %v", err)
			}
			return server.Shutdown(ctx)
	}
}

