package marker

import (
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
)

var Save = map[string]func(img *image.RGBA, path string){
	".png":  SavePNG,
	".jpeg": SaveJPEG,
	".jpg":  SaveJPEG,
	".JPG":  SaveJPEG,
	"JPEG":  SaveJPEG,
	"auto":  AutoSave,
}

func AutoSave(img *image.RGBA, path string) {
	ext := filepath.Ext(path)
	if ext == ".png" {
		SavePNG(img, path)
	} else {
		SaveJPEG(img, path)
	}
}

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
