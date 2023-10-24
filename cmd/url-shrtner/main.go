package main

import (
	"log/slog"
	"os"
	"url-shrtnr/internal/config"
	"url-shrtnr/internal/lib/logger/sl"
	"url-shrtnr/internal/storage/pgsql"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info("Starting url-shrtnr", slog.String("env", cfg.Env))
	log.Debug("Debug messages are enabled")

	log.Info("Storage path", slog.String("env", cfg.StoragePath))
	storage, err := pgsql.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to init database storage", sl.Err(err))
		os.Exit(1)
	}

	alias, urlToSave := "", "https://yandex.ru/search/?text=parenthesis&lr=213&clid=2437996"

	id, err := storage.SaveURL(urlToSave, alias)
	if err != nil {
		log.Error("failed to save url in storage", sl.Err(err))
	}

	log.Info("Saved URL UUID", slog.String("uuid", id))

	//TODO:init router: chi, chi render | net-http

	//TODO:run server:
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
