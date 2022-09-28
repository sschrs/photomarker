package marker

import (
	"log"
	"path/filepath"
)

func Glob(path string) []string {
	files, err := filepath.Glob(path)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return files
}
