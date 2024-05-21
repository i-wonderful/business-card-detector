package tesseract

import (
	"log"
	"os/exec"
	"strings"
)

const ENG = "eng"
const RUS = "rus"
const SCRIPT_CYR = "Cyrillic"

func DetectLang(path string) string {
	cmd := exec.Command("tesseract", "--psm", "0", path, "stdout")

	output, err := cmd.Output()
	if err != nil {
		log.Printf("Error: %v", err)
		return ""
	}

	script := findScript(output)
	if script == SCRIPT_CYR {
		return RUS
	} else {
		return ENG
	}
}

func findScript(output []byte) string {
	script := ""

	outputStr := string(output)
	lines := strings.Split(outputStr, "\n")
	for _, line := range lines {
		if strings.Contains(line, "Script:") {
			return strings.TrimSpace(strings.TrimPrefix(line, "Script:"))
		}
	}
	return script
}
