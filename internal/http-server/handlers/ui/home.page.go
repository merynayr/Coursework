package ui

import (
	"Coursework/internal/lib/logger/sl"
	"html/template"
	"log/slog"
	"net/http"
)

// Обработчик главной страницы.
func New(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		// Инициализируем срез содержащий пути к двум файлам. Обратите внимание, что
		// файл home.page.tmpl должен быть *первым* файлом в срезе.
		files := []string{
			"./ui/html/home.page.html",
		}

		ts, err := template.ParseFiles(files...)
		if err != nil {
			log.Error("failed to ParseFiles", sl.Err(err))
			http.Error(w, "Internal Server Error", 500)
			return
		}

		err = ts.Execute(w, nil)
		if err != nil {
			log.Error("failed to Execute", sl.Err(err))
			http.Error(w, "Internal Server Error", 500)
		}
	}
}
