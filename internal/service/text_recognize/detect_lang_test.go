package text_recognize

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDetectLang(t *testing.T) {
	pathRus := "/home/olga/projects/card_detector/integration_test/testdata/rus_card.jpeg"
	resRus := DetectLang(pathRus)
	assert.Equal(t, "rus", resRus)

	pathEng := "/home/olga/projects/card_detector/integration_test/testdata/eng_card2.jpeg"
	resEng := DetectLang(pathEng)
	assert.Equal(t, "eng", resEng)
}
