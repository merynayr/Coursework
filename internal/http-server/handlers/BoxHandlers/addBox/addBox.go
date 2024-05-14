package addBox

import (
	"errors"
	"io"
	"net/http"

	"Coursework/internal/lib/logger/sl"
	"Coursework/internal/storage"
	"Coursework/internal/storage/sqlite"

	resp "Coursework/internal/lib/api/response"

	"log/slog"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

type Request struct {
	BoxID  int     `json:"box_id" validate:"required"`
	Status string  `json:"status" validate:"required"`
	Floor  int     `json:"floor" validate:"required"`
	Area   float64 `json:"area" validate:"required"`
}

type Response struct {
	resp.Response
}

type BoxSaver interface {
	AddBox(box sqlite.Box) (int64, error)
}

func New(log *slog.Logger, boxSaver BoxSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.box.save.New"

		log = log.With(slog.String("op", op))

		var req Request
		err := render.DecodeJSON(r.Body, &req)

		if errors.Is(err, io.EOF) {
			// Такую ошибку встретим, если получили запрос с пустым телом.
			// Обработаем её отдельно
			log.Error("request body is empty")

			render.JSON(w, r, resp.Error("empty request"))

			return
		}
		if err != nil {
			log.Error("failed to decode request body", sl.Err(err))

			render.JSON(w, r, resp.Error("failed to decode request"))

			return
		}

		log.Info("request body decoded", slog.Any("request", req))

		if err := validator.New().Struct(req); err != nil {
			validateErr := err.(validator.ValidationErrors)

			log.Error("invalid request", sl.Err(err))

			render.JSON(w, r, resp.ValidationError(validateErr))

			return
		}

		box := sqlite.Box{BoxID: req.BoxID, Status: req.Status, Floor: req.Floor, Area: req.Area}
		id, err := boxSaver.AddBox(box)
		if errors.Is(err, storage.ErrExists) {
			// log.Info("box already exists", slog.String("box", req.Name))
			log.Info("box already exists")
			render.JSON(w, r, resp.Error("box already exists"))

			return
		}
		if err != nil {
			log.Error("failed to add box", sl.Err(err))

			render.JSON(w, r, resp.Error("failed to add box"))

			return
		}
		log.Info("box added", slog.Int64("id", id))

		responseOK(w, r)
	}
}

func responseOK(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, Response{
		Response: resp.OK(),
	})
}
