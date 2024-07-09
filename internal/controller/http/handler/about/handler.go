package about

import (
	"html/template"
	"log"
	"net/http"

	. "card_detector/internal/controller/http/data"
)

type Handler struct {
	version string
}

func NewAboutHandler(version string) *Handler {
	return &Handler{
		version: version,
	}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	data := ProjectInfo{
		Version: h.version,
	}

	tmpl, err := template.ParseFiles("./template/about.html")
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}
}
