package model

type Image struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type Prediction struct {
	X          float64 `json:"x"`
	Y          float64 `json:"y"`
	Width      float64 `json:"width"`
	Height     float64 `json:"height"`
	Confidence float64 `json:"confidence"`
	Class      string  `json:"class"`
	ClassID    int     `json:"class_id"`
}

type ApiResponse struct {
	Time        float64      `json:"time"`
	Image       Image        `json:"image"`
	Predictions []Prediction `json:"predictions"`
}
