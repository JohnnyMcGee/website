package main

import (
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/JohnnyMcGee/website/api/contact"
	"github.com/JohnnyMcGee/website/api/handler"
	"github.com/joho/godotenv"
)

type ContactResult struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error,omitempty"`
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found: %v", err)
	}
	config, err := NewConfig(os.Getenv)
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	if err := Run(
		config,
		os.OpenFile,
		os.Stdout,
	); err != nil {
		log.Fatalf("server error: %v", err)
	}
}

func Run(
	config Config,
	openFile func(name string, flag int, perm os.FileMode) (*os.File, error),
	stdout io.Writer,
) error {
	srv := http.NewServeMux()
	var logger *slog.Logger

	f, err := openFile(config.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return fmt.Errorf("error opening log file: %w", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Printf("error closing log file: %v", err)
		}
	}()

	logger = slog.New(slog.NewJSONHandler(f, nil))
	logger = logger.With("version", config.Version)
	logger.Info("ServerStarted", "host", config.Host, "port", config.Port)
	contactApi := contact.New(logger)
	contactHandler := handler.NewContactHandler(contactApi, logger)

	srv.HandleFunc("/contact", contactHandler.SendMessage)

	if err := http.ListenAndServe(":8000", srv); err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}
