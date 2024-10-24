package router

import (
	"net/http"

	"card_detector/internal/controller/http/handler/about"
	"card_detector/internal/controller/http/handler/file_upload"
	"card_detector/internal/controller/http/handler/file_upload_ui"
	"card_detector/internal/controller/http/handler/history"
	"card_detector/internal/controller/http/handler/index"
	"card_detector/pkg/log"
)

func NewRouter(detectService file_upload.Detector, getterService history.Getter, tmpFilePath, version string, logger *log.Logger) *http.ServeMux {
	// Создаем файловый сервер, который будет использовать директорию "./template/static"
	fs := http.FileServer(http.Dir("./template/static"))

	// Создаем ServeMux, который будет маршрутизировать запросы
	mux := http.NewServeMux()

	// Правильно указываем паттерн и убираем префикс
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	fsStorage := http.FileServer(http.Dir("./storage"))
	mux.Handle("/storage/", http.StripPrefix("/storage/", fsStorage))

	indexHandler := index.NewIndexHandler(version, logger)
	detectHandler := file_upload.NewFileUploadHandler(detectService, tmpFilePath, logger)
	detectHandlerUI := file_upload_ui.NewHandler(detectService, tmpFilePath)
	historyHandler := history.NewHandler(getterService, logger)
	aboutHandler := about.NewAboutHandler(version, logger)

	mux.HandleFunc("/", indexHandler.Handle)
	mux.HandleFunc("/detect", detectHandler.Handle)
	mux.HandleFunc("/detect_ui", detectHandlerUI.Handle)
	mux.HandleFunc("/history", historyHandler.Handle)
	mux.HandleFunc("/about", aboutHandler.Handle)

	return mux
}
