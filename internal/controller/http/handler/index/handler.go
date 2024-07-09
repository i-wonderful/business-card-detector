package index

import (
	data2 "card_detector/internal/controller/http/data"
	"html/template"
	"log"
	"net/http"
)

type IndexHandler struct {
	version string
}

func NewIndexHandler(version string) *IndexHandler {
	return &IndexHandler{
		version: version,
	}
}

func (h *IndexHandler) Handle(w http.ResponseWriter, r *http.Request) {
	// Создаем объект данных для шаблона
	data := data2.ProjectInfo{
		Version: h.version,
	}

	// Парсим шаблон из файла
	tmpl, err := template.ParseFiles("./template/index.html")
	if err != nil {
		log.Println(err)
	}

	// Генерируем вывод на основе шаблона и данных
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
	}
}
