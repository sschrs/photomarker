package marker

var usages = map[string]string{
	"width":   "Specifies the width, in pixels, of the photo to be placed. Default: 60px",
	"height":  "Specifies the height, in pixels, of the photo to be placed. Default: 60px",
	"opacity": "Opacity value of the photo to be placed. Min: 0, Max:1, Default: 0.5",
	"format":  "Specifies the output format of the marked photo. If the value of auto is set, the format of the highlighted photo is used as the output format. Possible values: png, jpg, jpeg, auto. default: auto",
	"replace": "Indicates whether to replace the photo or create a copy.",
	"out":     "If the replace argument is false, it specifies the destination path of the copied photos. Default: marked/",
	"x":       "Specifies the x coordinate of the photo to be placed. If a negative value is entered, pixels from the lower right corner are counted. Default: -20",
	"y":       "Specifies the y coordinate of the photo to be placed. If a negative value is entered, pixels from the lower right corner are counted. Default: -20",
}
