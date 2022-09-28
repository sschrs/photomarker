package marker

import (
	"log"
	"os"
)

type Marker struct {
	width, height, x, y       int
	opacity                   float64
	format, srcPhoto, outPath string
	dstPhotos                 []string
	replace                   bool
}

func NewMarker(width, height, x, y int, opacity float64, format string, replace bool) *Marker {
	return &Marker{
		width:   width,
		height:  height,
		x:       x,
		y:       y,
		opacity: opacity,
		format:  format,
		replace: replace,
	}
}

func (marker *Marker) AddFile(path string) {
	marker.dstPhotos = append(marker.dstPhotos, path)
}

func (marker *Marker) SetSrcFile(path string) {
	stat, err := os.Stat(path)
	if err != nil {
		log.Fatalln(err.Error())
	}
	if !stat.IsDir() && ExtCheck(path) {
		marker.srcPhoto = path
	} else {
		log.Fatalln("Marker photo must be png, jpg or jpeg.")
	}
}
