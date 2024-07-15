package integration_test

//import (
//	"bytes"
//	"card_detector/internal/model"
//	"card_detector/internal/rest"
//	"encoding/json"
//	"github.com/stretchr/testify/assert"
//	"io"
//	"mime/multipart"
//	"net/http"
//	"net/http/httptest"
//	"os"
//	"path/filepath"
//	"testing"
//)
//
//const UPLOAD_URL = "/upload"

//func TestDetectCards(t *testing.T) {
//	ts := httptest.NewServer(rest.NewServer().Handler)
//	defer ts.Close()
//
//	for _, tc := range testCases {
//		t.Run(tc.name, func(t *testing.T) {
//			runUploadFileIntegrationTest(t, tc, ts.URL+UPLOAD_URL)
//		})
//	}
//}

//func runUploadFileIntegrationTest(t *testing.T, tc testCase, url string) {
//	file, err := os.Open(tc.filePath)
//	assert.NoError(t, err, "could not open file")
//	defer file.Close()
//
//	body := &bytes.Buffer{}
//	writer := multipart.NewWriter(body)
//	part, err := writer.CreateFormFile("File", filepath.Base(tc.filePath))
//	assert.NoError(t, err, "could not create form file")
//
//	io.Copy(part, file)
//	writer.Close()
//
//	req, err := http.NewRequest("POST", url, body)
//	assert.NoError(t, err, "could not create request")
//
//	req.Header.Set("Content-Type", writer.FormDataContentType())
//	resp, err := http.DefaultClient.Do(req)
//	if err != nil {
//		t.Fatalf("failed to make request: %v", err)
//	}
//	defer resp.Body.Close()
//
//	assert.Equal(t, http.StatusOK, resp.StatusCode, "Status code")
//
//	var actual model.Person
//	err = json.NewDecoder(resp.Body).Decode(&actual)
//	assert.NoError(t, err, "could not decode response")
//
//	assert.Equal(t, tc.expected.Name, actual.Name, "Name")
//	assert.Equal(t, tc.expected.Email, actual.Email, "Email")
//	assert.Equal(t, tc.expected.Phone, actual.Phone, "Phone")
//	assert.Equal(t, tc.expected.JobTitle, actual.JobTitle, "JobTitle")
//	assert.Equal(t, tc.expected.Telegram, actual.Telegram, "Telegram")
//	assert.Equal(t, tc.expected.Site, actual.Site, "Site")
//	assert.Equal(t, tc.expected.Skype, actual.Skype, "Skype")
//	assert.Equal(t, tc.expected.Organization, actual.Organization, "Organization")
//}
