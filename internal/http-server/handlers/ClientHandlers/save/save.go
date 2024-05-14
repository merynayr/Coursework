package save

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
	ClientID int    `json:"clientID" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Type     string `json:"type" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
}

type Response struct {
	resp.Response
}

type ClientSaver interface {
	AddClient(client sqlite.Client) (int64, error)
}

func New(log *slog.Logger, cltSaver ClientSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.client.save.New"

		log = log.With(slog.String("op", op))
		println(r.Body)
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

		clnt := sqlite.Client{ClientID: req.ClientID, Name: req.Name, Type: req.Type, Phone: req.Phone}
		id, err := cltSaver.AddClient(clnt)
		if errors.Is(err, storage.ErrExists) {
			log.Info("client already exists", slog.String("client", req.Name))

			render.JSON(w, r, resp.Error("client already exists"))

			return
		}
		if err != nil {
			log.Error("failed to add client", sl.Err(err))

			render.JSON(w, r, resp.Error("failed to add client"))

			return
		}
		log.Info("client added", slog.Int64("id", id))

		responseOK(w, r)
	}
}

func responseOK(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, Response{
		Response: resp.OK(),
	})
}
