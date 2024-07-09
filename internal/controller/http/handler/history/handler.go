package history

import (
	"card_detector/internal/model"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Getter interface {
	GetAll() []model.Card
}

type Handler struct {
	getter Getter
}

func NewHandler(getter Getter) *Handler {
	return &Handler{
		getter: getter,
	}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{
		"formatDate": formatDate,
	}

	tmpl, err := template.New("history.html").
		Funcs(funcMap).
		ParseFiles("./template/history.html")
	if err != nil {
		log.Println(err)
	}

	cards := h.getter.GetAll()
	//data := []interface{}{cards}

	// Генерируем вывод на основе шаблона и данных
	err = tmpl.Execute(w, cards)
	if err != nil {
		log.Println(err)
	}
}

// Функция для форматирования времени
func formatDate(t time.Time) string {
	return t.Format("02.01.2006 15:04")
}
