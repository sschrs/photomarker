package marker

import (
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
)

func PNGToImg(path string) image.Image {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return img
}

func JPGToImg(path string) image.Image {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer file.Close()

	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatalln(err)
	}
	return img
}
