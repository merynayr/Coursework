package selectBox

import (
	"Coursework/internal/lib/logger/sl"
	"Coursework/internal/storage/sqlite"
	"encoding/json"
	"net/http"

	"log/slog"
)

type BoxGetter interface {
	SelectBoxes() ([]sqlite.Box, error)
}

func Select(log *slog.Logger, boxGetter BoxGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.box.select"
		log = log.With(slog.String("op", op))

		boxes, err := boxGetter.SelectBoxes() // Получаем клиентов из базы данных
		if err != nil {
			log.Error("Failed to get boxes", sl.Err(err))
			return
		}
		println(boxes)
		// Отправляем список клиентов в формате JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(boxes)
	}
}
