module card_detector

go 1.20

require (
	// img
	github.com/aaronland/go-image v1.2.3
	github.com/disintegration/imaging v1.6.2
	github.com/google/uuid v1.6.0
	github.com/nfnt/resize v0.0.0-20180221191011-83c6a9932646

	// detect
	github.com/rwcarlsen/goexif v0.0.0-20190401172101-9e8deecbddbd

	// util
	github.com/stretchr/testify v1.8.4
	golang.org/x/image v0.16.0
	gopkg.in/yaml.v3 v3.0.1
)

require github.com/yalue/onnxruntime_go v1.11.0

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
)
