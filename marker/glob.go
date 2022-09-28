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

func ExtCheck(path string) bool {
	ext := filepath.Ext(path)
	if ext == ".png" || ext == ".jpg" || ext == ".jpeg" {
		return true
	}
	return false
}
