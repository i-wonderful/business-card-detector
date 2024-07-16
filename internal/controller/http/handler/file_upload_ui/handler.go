package file_upload_ui

import (
	"card_detector/internal/model"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type Response struct {
	Person  *model.Person `json:"person"`
	ImgPath string        `json:"img_path"`
}

type Detector interface {
	Detect(pathImg string) (*model.Person, string, error)
}

type Handler struct {
	name       string
	detector   Detector
	dirTmpPath string
}

func NewHandler(detector Detector, dirTmpPath string) *Handler {
	return &Handler{
		name:       "FileUploadHandler",
		detector:   detector,
		dirTmpPath: dirTmpPath,
	}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
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
		log.Println(err)
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

	person, filePath, err := h.detector.Detect(fileName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&Response{person, filePath})

	duration := time.Since(start)
	fmt.Printf(">>> [Time] %s took %v\n", "Full detection", duration)
}
