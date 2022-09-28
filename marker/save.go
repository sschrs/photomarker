package marker

import (
	"image"
	"image/png"
	"log"
	"os"
)

func SavePNG(img *image.RGBA, path string) {
	f, err := os.Create(path)
	if err != nil {
		log.Fatalln(err.Error())
	}
	png.Encode(f, img)
}
