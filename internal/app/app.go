package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"card_detector/internal/controller/http/router"
	"card_detector/internal/repo/inmemory"
	"card_detector/internal/service"
	"card_detector/internal/service/detect/onnx"
	"card_detector/internal/service/field_sort"
	shistory "card_detector/internal/service/history"
	"card_detector/internal/service/img_prepare"
	"card_detector/internal/service/text_recognize/paddleocr"

	"card_detector/pkg/log"
)

type app struct {
	config *Config
}

func NewApp2(config *Config) *app {
	return &app{
		config: config,
	}
}

func (a *app) Run() error {

	logger, err := log.NewLogger("app", a.config.Log.Level, false)

	// repo
	cardRepo := inmemory.NewCardRepo()

	// service
	imgPreparer := img_prepare.NewService(a.config.StorageFolder, a.config.TmpFolder, logger)

	isLogTime := a.config.Log.Time
	textRecognizer, err := paddleocr.NewService(isLogTime,
		a.config.Paddleocr.RunPath,
		a.config.Paddleocr.DetPath,
		a.config.Paddleocr.RecPath,
		a.config.TmpFolder,
		logger)
	if err != nil {
		logger.Fatal("text recognizer creation error", err)
	}
	cardDetector, err := onnx.NewService(
		a.config.Onnx.PathRuntime,
		a.config.Onnx.PathModel,
		isLogTime,
		logger)
	if err != nil {
		logger.Fatal("card detector creation error", err)
	}

	fieldSorter := field_sort.NewService(
		a.config.PathProfessionList,
		a.config.PathCompanyList,
		a.config.PathNamesList,
		isLogTime,
		logger)
	getterService := shistory.NewService(cardRepo)

	// detector
	detectService := service.NewDetector2(
		imgPreparer,
		textRecognizer,
		cardDetector,
		fieldSorter,
		cardRepo,
		a.config.StorageFolder,
		a.config.TmpFolder,
		isLogTime,
		logger)

	// handlers
	h := router.NewRouter(detectService, getterService, a.config.TmpFolder, a.config.Version, logger)

	// start server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", a.config.Port),
		Handler: h,
	}
	go func() {
		logger.Info(fmt.Sprintf("Starting app: %s version: %s", a.config.Name, a.config.Version))
		logger.Info(fmt.Sprintf("Listening on port: %d", a.config.Port))
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			logger.Fatal("ListenAndServe()", err)
		}
	}()

	// Set up a channel to listen for an interrupt or terminate signal from the OS.
	// Use a buffered channel to prevent missing signals.
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Block until a signal is received.
	<-stop

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait until the timeout deadline.
	logger.Info("Shutting down server...")
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Server Shutdown", err)
	}

	logger.Info("Server gracefully stopped")
	return nil
}
