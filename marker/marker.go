package marker

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
