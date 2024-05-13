package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"card_detector/internal/controller/http/router"
	"card_detector/internal/repo/inmemory"
	"card_detector/internal/service"
	"card_detector/internal/service/field_sort"
	shistory "card_detector/internal/service/history"
	"card_detector/internal/service/img_prepare"
	"card_detector/internal/service/text_find/onnx"
	"card_detector/internal/service/text_recognize"
)

type app struct {
	config *AppConfig
}

func NewApp(config *AppConfig) *app {
	return &app{
		config: config,
	}
}

func (a *app) Run() error {

	// repo
	cardRepo := inmemory.NewCardRepo()

	// service
	imgPreparer := img_prepare.NewService(a.config.StorageFolder)

	isLogTime := a.config.Log.Time
	findTextService, err := onnx.NewService(a.config.Onnx.PathRuntime, a.config.Onnx.PathModel, isLogTime) //findTextService := remote.NewService()
	if err != nil {
		log.Fatal(err)
		return err
	}
	textRecognizer := text_recognize.NewService(isLogTime, "./config/tesseract/")
	fieldSorter := field_sort.NewService(
		a.config.PathProfessionList,
		a.config.PathCompanyList,
		a.config.PathNamesList,
		isLogTime)
	getterService := shistory.NewService(cardRepo)

	// detector
	detectService := service.NewDetector(
		imgPreparer,
		findTextService,
		textRecognizer,
		fieldSorter,
		cardRepo,
		a.config.StorageFolder,
		isLogTime)

	// handlers
	h := router.NewRouter(detectService, getterService)

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
