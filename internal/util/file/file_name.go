package manage_file

import (
	"github.com/google/uuid"
	"strings"
)

func GenerateFileName(pathStorage, prefix, format string) string {
	sb := strings.Builder{}
	sb.WriteString(pathStorage)
	sb.WriteString("/")
	if prefix != "" {
		sb.WriteString(prefix)
		sb.WriteString("_")
	}
	sb.WriteString(uuid.New().String())
	sb.WriteString(".")
	sb.WriteString(format)
	return sb.String()
}
