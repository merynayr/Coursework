package main

import (
	"fmt"
	"log/slog"
	"os"

	"Coursework/internal/config"
	"Coursework/internal/storage/sqlite"
)

func main() {
	// TODO: init config

	cfg := config.MustLoad()
	fmt.Println((cfg))

	// TODO: init logger

	log := setupLogger(cfg.Env)

	log.Info("starting Project", slog.String("env", cfg.Env))

	//TODO: init storage

	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to init storage", err)
		os.Exit(1)
	}

	_ = storage

	log.Info("starting storage")

	//TODO: init router

	//TODO: run server
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case "local":
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	}

	return log
}
