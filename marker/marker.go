package marker

import (
	"fmt"
	"golang.org/x/image/draw"
	"image"
	"image/color"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type Marker struct {
	width, height, x, y       int
	opacity                   float64
	format, srcPhoto, outPath string
	dstPhotos                 []string
	replace                   bool
}

func NewMarker(width, height, x, y int, opacity float64, format, out string, replace bool) *Marker {
	return &Marker{
		width:   width,
		height:  height,
		x:       x,
		y:       y,
		opacity: opacity,
		format:  format,
		replace: replace,
		outPath: out,
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

func (marker *Marker) Merge() {
	start := time.Now()

	srcImg := Resize(marker.width, marker.height, marker.srcPhoto)
	limit := make(chan struct{}, 12)
	var wg sync.WaitGroup
	for _, photo := range marker.dstPhotos {
		wg.Add(1)
		go func(photo string) {
			limit <- struct{}{}
			defer func() { <-limit }()
			defer wg.Done()
			marker.Draw(srcImg, photo)
		}(photo)
	}
	wg.Wait()
	close(limit)
	end := time.Now()
	fmt.Println(fmt.Sprintf("%v for %d files", end.Sub(start), len(marker.dstPhotos)))
}

func (marker *Marker) Draw(srcImg *image.RGBA, photo string) {
	img := GetImage(photo)

	x := marker.x
	y := marker.y

	if x < 0 {
		x = img.Bounds().Size().X + x - srcImg.Bounds().Size().X
	}

	if y < 0 {
		y = img.Bounds().Size().Y + y - srcImg.Bounds().Size().Y
	}

	mask := image.NewUniform(color.Alpha{uint8(float64(255) * marker.opacity)})
	newImage := image.NewRGBA(image.Rect(0, 0, img.Bounds().Size().X, img.Bounds().Size().Y))

	draw.Draw(newImage, newImage.Bounds(), img, image.Point{0, 0}, draw.Src)
	draw.DrawMask(newImage, newImage.Bounds(), srcImg, image.Point{-x, -y}, mask, image.Point{-x, -y}, draw.Over)

	dirPath := filepath.Dir(photo)
	os.Chdir(dirPath)
	if err := os.MkdirAll(fmt.Sprintf(marker.outPath), 0777); err != nil {
		log.Fatalln(err.Error())
	}
	fileName := filepath.Base(photo)
	Save[marker.format](newImage, fmt.Sprintf("%s/%s", marker.outPath, fileName))
	fmt.Println(fmt.Sprintf("%s marked", fileName))

}
