package addContract

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
	ContractID int    `json:"ContractID" validate:"required"`
	ClientID   int    `json:"ClientID" validate:"required"`
	BoxID      int    `json:"BoxID" validate:"required"`
	DateSigned string `json:"DateSigned" validate:"required"`
	StartDate  string `json:"StartDate" validate:"required"`
	EndDate    string `json:"EndDate" validate:"required"`
}

type Response struct {
	resp.Response
}

type ContractSaver interface {
	AddContract(contract sqlite.Contract) (int64, error)
}

func Add(log *slog.Logger, ContractSaver ContractSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.contract.save.Add"

		log = log.With(slog.String("op", op))

		var req Request
		err := render.DecodeJSON(r.Body, &req)
		println(r.Body)
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

		contract := sqlite.Contract{ContractID: req.ContractID, ClientID: req.ClientID, BoxID: req.BoxID, DateSigned: req.DateSigned, StartDate: req.StartDate, EndDate: req.EndDate}
		id, err := ContractSaver.AddContract(contract)
		if errors.Is(err, storage.ErrExists) {
			// log.Info("contract already exists", slog.String("contract", req.Name))
			log.Info("contract already exists")
			render.JSON(w, r, resp.Error("contract already exists"))

			return
		}
		if err != nil {
			log.Error("failed to add contract", sl.Err(err))

			render.JSON(w, r, resp.Error("failed to add contract"))

			return
		}
		log.Info("contract added", slog.Int64("id", id))

		responseOK(w, r)
	}
}

func responseOK(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, Response{
		Response: resp.OK(),
	})
}
