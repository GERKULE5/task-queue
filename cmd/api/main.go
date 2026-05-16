package main

import (
	"log/slog"
	"os"
	"task_queue/pkg/logger"
	"task_queue/pkg/postgres"
)

func main() {
	err := logger.Setup()
	if err != nil {
		slog.Error(err.Error())
	}

	slog.Info("Server is starting...")

	db, err := postgres.NewConnection()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()
}
