package router

import (
	"card_detector/internal/controller/http/handler/about"
	"card_detector/internal/controller/http/handler/file_upload"
	"card_detector/internal/controller/http/handler/history"
	"card_detector/internal/controller/http/handler/index"
	"net/http"
)

func NewRouter(detectService file_upload.Detector, getterService history.Getter, tmpFilePath, version string) *http.ServeMux {
	// Создаем файловый сервер, который будет использовать директорию "./template/static"
	fs := http.FileServer(http.Dir("./template/static"))

	// Создаем ServeMux, который будет маршрутизировать запросы
	mux := http.NewServeMux()

	// Правильно указываем паттерн и убираем префикс
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	fsStorage := http.FileServer(http.Dir("./storage"))
	mux.Handle("/storage/", http.StripPrefix("/storage/", fsStorage))

	indexHandler := index.NewIndexHandler(version)
	detectHandler := file_upload.NewFileUploadHandler(detectService, tmpFilePath)
	historyHandler := history.NewHandler(getterService)
	aboutHandler := about.NewAboutHandler(version)

	mux.HandleFunc("/", indexHandler.Handle)
	mux.HandleFunc("/detect", detectHandler.Handle)
	mux.HandleFunc("/history", historyHandler.Handle)
	mux.HandleFunc("/about", aboutHandler.Handle)

	return mux
}
