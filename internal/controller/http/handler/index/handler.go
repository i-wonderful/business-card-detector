package index

import (
	"html/template"
	"log"
	"net/http"
)

// Создаем структуру данных для передачи в шаблон
type MyData struct {
	Title string
	Name  string
}

type IndexHandler struct {
}

func NewIndexHandler() *IndexHandler {
	return &IndexHandler{}
}

func (h *IndexHandler) Handle(w http.ResponseWriter, r *http.Request) {
	// Создаем объект данных для шаблона
	data := MyData{
		Title: "Пример веб-приложения",
		Name:  "Вася",
	}

	// Парсим шаблон из файла
	tmpl, err := template.ParseFiles("./template/index.html")
	if err != nil {
		log.Fatal(err)
	}

	// Генерируем вывод на основе шаблона и данных
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}
}
