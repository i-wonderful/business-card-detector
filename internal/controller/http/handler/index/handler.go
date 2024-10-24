package index

import (
	"html/template"
	"net/http"

	data2 "card_detector/internal/controller/http/data"
	"card_detector/pkg/log"
)

type Handler struct {
	version string
	log     *log.Logger
}

func NewIndexHandler(version string, log *log.Logger) *Handler {
	return &Handler{
		version: version,
		log:     log,
	}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	// Создаем объект данных для шаблона
	data := data2.ProjectInfo{
		Version: h.version,
	}

	// Парсим шаблон из файла
	tmpl, err := template.ParseFiles("./template/index.html")
	if err != nil {
		h.log.Error("Error parse template index.html", err)
	}

	// Генерируем вывод на основе шаблона и данных
	err = tmpl.Execute(w, data)
	if err != nil {
		h.log.Error("Error execute template index.html", err)
	}
}
