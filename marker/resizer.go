package marker

import (
	"golang.org/x/image/draw"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
)

func Resize(width, height int, imagePath string) *image.RGBA {
	ext := filepath.Ext(imagePath)
	file, err := os.Open(imagePath)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer file.Close()

	var decodeError error
	var img image.Image

	if ext == ".png" {
		img, decodeError = png.Decode(file)
	} else {
		img, decodeError = jpeg.Decode(file)
	}

	if decodeError != nil {
		log.Fatalln(decodeError.Error())
	}

	newImage := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.NearestNeighbor.Scale(newImage, newImage.Rect, img, img.Bounds(), draw.Over, nil)

	return newImage
}
