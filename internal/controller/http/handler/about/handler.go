package about

import (
	"html/template"
	"net/http"

	. "card_detector/internal/controller/http/data"
	"card_detector/pkg/log"
)

type Handler struct {
	version string
	logger  *log.Logger
}

func NewAboutHandler(version string, logger *log.Logger) *Handler {
	return &Handler{
		version: version,
		logger:  logger,
	}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	data := ProjectInfo{
		Version: h.version,
	}

	tmpl, err := template.ParseFiles("./template/about.html")
	if err != nil {
		h.logger.Fatal("Error parse template", err)
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		h.logger.Fatal("Error execute template", err)
	}
}
