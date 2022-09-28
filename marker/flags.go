package marker

import (
	"flag"
	"log"
)

var (
	width   = flag.Int("width", 100, usages["width"])
	height  = flag.Int("height", 100, usages["height"])
	x       = flag.Int("x", -20, usages["x"])
	y       = flag.Int("y", -20, usages["y"])
	opacity = flag.Float64("opacity", 0.5, usages["opacity"])
	format  = flag.String("format", "auto", usages["format"])
	replace = flag.Bool("replace", false, usages["replace"])
	out     = flag.String("out", "marked/", usages["out"])
	dst     = flag.String("dst", "", usages["dst"])
	src     = flag.String("src", "", usages["src"])
)

func CreateMarkerFromFlags() *Marker {
	flag.Parse()

	if *src == "" || *dst == "" {
		log.Fatalln("-dst and -src args are required")
	}

	m := NewMarker(*width, *height, *x, *y, *opacity, *format, *out, *replace)
	m.SetSrcFile(*src)
	AddFiles(m, *dst)
	return m
}
