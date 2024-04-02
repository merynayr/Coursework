package update

import (
	"errors"
	"io"
	"net/http"

	"Coursework/internal/lib/logger/sl"
	"Coursework/internal/storage/sqlite"

	resp "Coursework/internal/lib/api/response"

	"log/slog"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

type Request struct {
	ClientID int    `json:"clientID"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Phone    string `json:"phone"`
}

type Response struct {
	resp.Response
}

type ClientUpdater interface {
	UpdateClient(client sqlite.Client) error
}

func Update(log *slog.Logger, cltUpdater ClientUpdater) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.client.update.Update"

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

		clnt := sqlite.Client{ClientID: req.ClientID, Name: req.Name, Type: req.Phone, Phone: req.Type}
		err = cltUpdater.UpdateClient(clnt)
		if err != nil {
			log.Error("failed to update client", sl.Err(err))

			render.JSON(w, r, resp.Error("failed to update client"))

			return
		}

		responseOK(w, r)
	}
}

func responseOK(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, Response{
		Response: resp.OK(),
	})
}
