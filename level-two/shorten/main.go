package main

import (
	"log/slog"
	"net/http"
	"shorten/api"
	"time"
)

func main() {
	if err := run(); err != nil {
		slog.Error("failed to execute code", "error", err)
		return
	}
	slog.Info("all systems offline")
}

func run() error {
	db := make(map[string]string)

	handler := api.NewHandler(db)

	s := http.Server{
		ReadTimeout: 10 * time.Second,
		IdleTimeout: 1 * time.Minute,
		WriteTimeout: 10 * time.Second,
		Addr: ":8080",
		Handler: handler,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}