package remote

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	model2 "card_detector/internal/model"
	"card_detector/internal/service/detect/remote/model"
)

type FindTextService struct {
	apiUrl string
	log    *log.Logger
}

func NewFindTextService() *FindTextService {
	return &FindTextService{
		apiUrl: os.Getenv("API_URL"), // todo
		log:    log.Default(),
	}
}
func (s *FindTextService) PredictTextCoord(pathImg string) ([]model2.TextArea, error) {

	response, err := getPrediction3(pathImg)
	if err != nil {
		fmt.Println("Error getting prediction:", err)
		return nil, fmt.Errorf("Error getting prediction: %v", err)
	}
	areas := []model2.TextArea{}
	for _, prediction := range response.Predictions {
		if prediction.Class == "TEXT" {
			xStart := prediction.X - prediction.Width/2.0
			yStart := prediction.Y - prediction.Height/2.0

			area := &model2.TextArea{
				X:      int(xStart),
				Y:      int(yStart),
				Width:  int(prediction.Width),
				Height: int(prediction.Height),
			}
			areas = append(areas, *area)
		}
	}
	return areas, nil
}

//	This code reads the image file, encodes it as base64,
//
// and sends it as the request body instead of using the multipart/form-data approach.
func getPrediction3(imagePath string) (*model.ApiResponse, error) {
	//  confidence=40, overlap=30
	apiURL := "https://detect.roboflow.com/business_cards-obnkp/5?api_key=RV4vxlFbQGUj4mMMw3bO&confidence=6&overlap=59"

	// Read the image file
	imageBytes, err := os.ReadFile(imagePath)
	if err != nil {
		fmt.Println("Error reading the image:", err)
		return nil, err
	}

	// Encode the image as base64
	base64Image := base64.StdEncoding.EncodeToString(imageBytes)

	// Create an HTTP POST request
	req, err := http.NewRequest("POST", apiURL, strings.NewReader(base64Image))
	if err != nil {
		fmt.Println("Error creating the request:", err)
		return nil, err
	}

	// Set request headers and parameters
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	q := req.URL.Query()
	//q.Add("api_key", apiKey)
	req.URL.RawQuery = q.Encode()

	// Send the request
	client := &http.Client{}
	startReq := time.Now()
	resp, err := client.Do(req)
	endReq := time.Now()
	log.Printf("Request took %v", endReq.Sub(startReq))
	if err != nil {
		fmt.Println("Error sending the request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Request failed with status code: %d\n", resp.StatusCode)
		return nil, fmt.Errorf("Request failed with status code: %d", resp.StatusCode)
	}

	// Read the response
	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading the response:", err)
		return nil, err
	}

	// Parse the JSON response
	var apiResponse model.ApiResponse
	err = json.Unmarshal(responseData, &apiResponse)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return nil, fmt.Errorf("Error parsing JSON: %v", err)
	}
	return &apiResponse, nil
}
