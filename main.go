package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/takahiromitsui/go-task-manager/application"
)

func main() {
		app := application.NewApp()
		ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
		defer cancel()
		err := app.Start(ctx)
		if err != nil {
			fmt.Printf("Error starting the server: %v\n", err)
		}
}