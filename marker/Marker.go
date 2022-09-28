package marker

type Marker struct {
	width, height, x, y       int
	opacity                   float64
	format, srcPhoto, outPath string
	dstPhotos                 []string
	replace                   bool
}
