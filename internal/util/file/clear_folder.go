package manage_file

import (
	"os"
	"path/filepath"
)

func ClearFolder(dir string) error {

	files, err := filepath.Glob(filepath.Join(dir, "*"))
	if err != nil {
		return err
	}

	for _, file := range files {
		err = os.Remove(file)
		if err != nil {
			return err
		}
	}

	return nil
}
