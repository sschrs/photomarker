package marker

import (
	"log"
	"os"
	"path/filepath"
)

func AddFiles(marker *Marker, path string) {
	dstStat, err := os.Stat(path)
	if err != nil {
		log.Fatalln(err.Error())
	}
	if dstStat.IsDir() {
		files := Glob(path + "/*")
		for _, file := range files {
			if ExtCheck(file) {
				marker.AddFile(file)
			}
		}
	} else {
		if ExtCheck(path) {
			marker.AddFile(path)
		}
	}

}

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
