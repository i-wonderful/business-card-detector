package history

import (
	"html/template"
	"net/http"
	"time"

	"card_detector/internal/model"
	"card_detector/pkg/log"
)

type Getter interface {
	GetAll() []model.Card
}

type Handler struct {
	getter Getter
	log    *log.Logger
}

func NewHandler(getter Getter, log *log.Logger) *Handler {
	return &Handler{
		getter: getter,
		log:    log,
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
		h.log.Error("Error parse template history", err)
	}

	cards := h.getter.GetAll()
	//data := []interface{}{cards}

	// Генерируем вывод на основе шаблона и данных
	err = tmpl.Execute(w, cards)
	if err != nil {
		h.log.Error("Error execute template history", err)
	}
}

// Функция для форматирования времени
func formatDate(t time.Time) string {
	return t.Format("02.01.2006 15:04")
}
