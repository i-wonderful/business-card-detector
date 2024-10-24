package file_upload

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"card_detector/internal/model"
	"card_detector/pkg/log"
)

type Detector interface {
	Detect(pathImg string) (*model.Person, string, error)
}

type FileUploadHandler struct {
	name       string
	detector   Detector
	dirTmpPath string
	log        *log.Logger
}

func NewFileUploadHandler(detector Detector, dirTmpPath string, log *log.Logger) *FileUploadHandler {
	return &FileUploadHandler{
		name:       "FileUploadHandler",
		detector:   detector,
		dirTmpPath: dirTmpPath,
	}
}

func (h *FileUploadHandler) Handle(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	err := r.ParseMultipartForm(10 << 20) // 10MB max
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get file from form data
	file, handler, err := r.FormFile("image")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.log.Error("Error get file", err)
		return
	}
	defer file.Close()

	// Save the file to disk
	fileName := h.dirTmpPath + "/" + handler.Filename
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()

	io.Copy(f, file)

	person, _, err := h.detector.Detect(fileName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)

	duration := time.Since(start)
	fmt.Printf(">>> [Time] %s took %v\n", "Full detection", duration)
}
