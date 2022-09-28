package marker

import (
	"image"
	"log"
	"os"
)

func GetImage(path string) image.Image {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalln(err.Error())
	}
	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return img
}
