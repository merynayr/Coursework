package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"Coursework/internal/config"
	"Coursework/internal/http-server/handlers/BoxHandlers/selectBox"
	"Coursework/internal/http-server/handlers/BoxHandlers/addBox"

	"Coursework/internal/http-server/handlers/ClientHandlers/delete"
	"Coursework/internal/http-server/handlers/ClientHandlers/redirect"
	"Coursework/internal/http-server/handlers/ClientHandlers/save"
	"Coursework/internal/http-server/handlers/ClientHandlers/update"

	"Coursework/internal/http-server/handlers/ui"
	mwLogger "Coursework/internal/http-server/middleware/logger"
	"Coursework/internal/lib/logger/handlers/slogpretty"
	"Coursework/internal/lib/logger/sl"
	"Coursework/internal/storage/sqlite"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// TODO: init config

	cfg := config.MustLoad()
	fmt.Println((cfg))

	// TODO: init logger

	log := setupLogger(cfg.Env)

	log.Info(
		"starting Project",
		slog.String("env", cfg.Env),
		slog.String("version", "123"),
	)

	//TODO: init storage

	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}

	log.Info("starting storage")

	//TODO: init router

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(mwLogger.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.HandleFunc("/", ui.New(log))
	router.Post("/сlient/add", save.New(log, storage))
	router.Post("/client/del", delete.Del(log, storage))
	router.Get("/client/select", redirect.New(log, storage))
	router.Post("/сlient/update", update.Update(log, storage))
	router.Get("/box/select", selectBox.Select(log, storage))
	router.Post("/box/add", addBox.New(log, storage))


	fileServer := http.FileServer(http.Dir("./ui/static/"))
	router.Handle("/ui/static/*", http.StripPrefix("/ui/static", fileServer))

	//TODO: run server

	log.Info("starting server", slog.String("address", cfg.Address))

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Error("failed to start server")
	}

	log.Info("server stopped")

}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case "local":
		log = setupPrettySlog()
	default:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
