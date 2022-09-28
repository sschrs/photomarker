package marker

func Run() {
	m := CreateMarkerFromFlags()
	m.Merge()
}
