package application

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	// "os"
	"time"

	_ "github.com/lib/pq"
)


type App struct {
	router http.Handler
	db 	 *sql.DB
}

func connectToDB() *sql.DB {
	host := "localhost"
	port := 5432
	user := "postgres"
	password := "password"
	dbname := "go_task"

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)


	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		fmt.Printf("Error connecting to the database: %v\n", err)
	}

	return db

}

func NewApp() *App {
	app:= &App{
		db: connectToDB(),
	}
	app.loadRoutes()
	return app
}

func (app *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":8080",
		Handler: app.router,
	}

	err:= app.db.Ping()
	if err != nil {
		return fmt.Errorf("failed to ping the database: %v", err)
	}

	defer func() {
		if err := app.db.Close(); err != nil {
			fmt.Printf("Error closing the database: %v\n", err)
		}
	}()
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

