package delete

import (
	"errors"
	"io"
	"net/http"

	"Coursework/internal/lib/logger/sl"
	"Coursework/internal/storage"

	resp "Coursework/internal/lib/api/response"

	"log/slog"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

type Request struct {
	ClientID int `json:"clientID" validate:"required"`
}

type Response struct {
	resp.Response
}

type ClientDel interface {
	DeleteClient(clientId int) error
}

func Del(log *slog.Logger, cltDel ClientDel) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.client.delete.Del"

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

		err = cltDel.DeleteClient(req.ClientID)
		if errors.Is(err, storage.ErrExists) {
			log.Info("client already dont exists")

			render.JSON(w, r, resp.Error("client dont exists"))

			return
		}
		if err != nil {
			log.Error("failed to add client", sl.Err(err))

			render.JSON(w, r, resp.Error("failed to del client"))

			return
		}
		log.Info("client deleted")

		responseOK(w, r)
	}
}

func responseOK(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, Response{
		Response: resp.OK(),
	})
}
