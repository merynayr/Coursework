package selectContract

import (
	"Coursework/internal/lib/logger/sl"
	"Coursework/internal/storage/sqlite"
	"encoding/json"
	"net/http"

	"log/slog"
)

type ContractGetter interface {
	SelectContracts() ([]sqlite.Contract, error)
}

func Select(log *slog.Logger, contractGetter ContractGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.contract.select"
		log = log.With(slog.String("op", op))

		contracts, err := contractGetter.SelectContracts() // Получаем клиентов из базы данных
		if err != nil {
			log.Error("Failed to get contracts", sl.Err(err))
			return
		}
		// Отправляем список клиентов в формате JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(contracts)
	}
}
