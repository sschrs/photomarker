package marker

import (
	"flag"
)

var (
	width   = flag.Int("width", 60, usages["width"])
	height  = flag.Int("height", 60, usages["height"])
	x       = flag.Int("x", -20, usages["x"])
	y       = flag.Int("y", -20, usages["y"])
	opacity = flag.Float64("opacity", 0.5, usages["opacity"])
	format  = flag.String("format", "auto", usages["format"])
	replace = flag.Bool("replace", false, usages["replace"])
	out     = flag.String("out", "marked/", usages["out"])
)

func CreateMarkerFromFlags() *Marker {
	flag.Parse()
	return &Marker{
		width:   *width,
		height:  *height,
		x:       *x,
		y:       *y,
		opacity: *opacity,
		format:  *format,
		replace: *replace,
		outPath: *out,
	}
}
