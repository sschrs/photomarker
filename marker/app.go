package marker

import "log"

func Run(args []string) {
	if len(args) < 3 {
		log.Fatalln("destination and source path must be given")
	}
	dstArg := args[1]
	srcArg := args[2]

	mrkr := CreateMarkerFromFlags()
	mrkr.SetSrcFile(srcArg)
	AddFiles(mrkr, dstArg)
}
