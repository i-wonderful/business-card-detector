module card_detector

go 1.20

require (
	// img
	github.com/aaronland/go-image v1.2.3
	github.com/disintegration/imaging v1.6.2
	github.com/nfnt/resize v0.0.0-20180221191011-83c6a9932646
	github.com/rwcarlsen/goexif v0.0.0-20190401172101-9e8deecbddbd

	// detect
	github.com/otiai10/gosseract v2.2.1+incompatible
	github.com/yalue/onnxruntime_go v1.6.0

	// util
	github.com/stretchr/testify v1.8.4
	gopkg.in/yaml.v3 v3.0.1
	github.com/google/uuid v1.5.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/otiai10/mint v1.6.3 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/image v0.0.0-20191009234506-e7c1f5e7dbb8 // indirect
)
