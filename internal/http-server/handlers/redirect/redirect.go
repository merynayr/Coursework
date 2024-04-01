package redirect

import (
	"Coursework/internal/storage/sqlite"
	"encoding/json"
	"fmt"
	"net/http"

	"log/slog"
)

type ClientGetter interface {
	SelectClients() ([]sqlite.Client, error)
}

func New(log *slog.Logger, cltGetter ClientGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		clients, err := cltGetter.SelectClients() // Получаем клиентов из базы данных
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to get clients: %v", err), http.StatusInternalServerError)
			return
		}

		// Отправляем список клиентов в формате JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(clients)
	}
}
