package marker

import (
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
)

func SavePNG(img *image.RGBA, path string) {
	f, err := os.Create(path)
	if err != nil {
		log.Fatalln(err.Error())
	}
	if err = png.Encode(f, img); err != nil {
		log.Fatalln(err.Error())
	}
}

func SaveJPEG(img *image.RGBA, path string) {
	f, err := os.Create(path)
	if err != nil {
		log.Fatalln(err.Error())
	}
	if err = jpeg.Encode(f, img, nil); err != nil {
		log.Fatalln(err.Error())
	}
}
