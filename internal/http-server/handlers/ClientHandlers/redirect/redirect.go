package redirect

import (
	"Coursework/internal/lib/logger/sl"
	"Coursework/internal/storage/sqlite"
	"encoding/json"
	"net/http"

	"log/slog"
)

type ClientGetter interface {
	SelectClients() ([]sqlite.Client, error)
}

func New(log *slog.Logger, cltGetter ClientGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.client.redirect.New"
		log = log.With(slog.String("op", op))

		clients, err := cltGetter.SelectClients() // Получаем клиентов из базы данных
		if err != nil {
			log.Error("Failed to get clients", sl.Err(err))
			return
		}

		// Отправляем список клиентов в формате JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(clients)
	}
}
