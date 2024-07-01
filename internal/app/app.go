package app

import (
	"card_detector/internal/controller/http/router"
	"card_detector/internal/repo/inmemory"
	"card_detector/internal/service"
	"card_detector/internal/service/detect/onnx"
	"card_detector/internal/service/field_sort"
	shistory "card_detector/internal/service/history"
	"card_detector/internal/service/img_prepare"
	"card_detector/internal/service/text_recognize/paddleocr"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type app2 struct {
	config *Config
}

func NewApp2(config *Config) *app2 {
	return &app2{
		config: config,
	}
}

func (a *app2) Run() error {

	// repo
	cardRepo := inmemory.NewCardRepo()

	// service
	imgPreparer := img_prepare.NewService(a.config.StorageFolder)

	isLogTime := a.config.Log.Time
	textRecognizer, err := paddleocr.NewService(isLogTime,
		a.config.Paddleocr.RunPath,
		a.config.Paddleocr.DetPath,
		a.config.Paddleocr.RecPath)
	if err != nil {
		log.Fatal("text recognizer creation error", err)
	}
	cardDetector, err := onnx.NewService(
		a.config.Onnx.PathRuntime,
		a.config.Onnx.PathModel,
		isLogTime)
	if err != nil {
		log.Fatal("card detector creation error", err)
	}

	fieldSorter := field_sort.NewService(
		a.config.PathProfessionList,
		a.config.PathCompanyList,
		a.config.PathNamesList,
		isLogTime)
	getterService := shistory.NewService(cardRepo)

	// detector
	detectService := service.NewDetector2(
		imgPreparer,
		textRecognizer,
		cardDetector,
		fieldSorter,
		cardRepo,
		a.config.StorageFolder,
		isLogTime,
		a.config.IsDebug)

	// handlers
	h := router.NewRouter(detectService, getterService, a.config.StorageFolder, a.config.Version)

	// start server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", a.config.Port),
		Handler: h,
	}
	go func() {
		log.Println("Starting app:", a.config.Name, a.config.Version)
		log.Println("Listening on port", a.config.Port)
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %v", err)
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
	log.Printf("Shutting down server...")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown: %v", err)
	}

	log.Printf("Server gracefully stopped")
	return nil
}
