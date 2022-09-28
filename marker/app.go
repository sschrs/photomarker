package marker

import "fmt"

func Run() {
	m := CreateMarkerFromFlags()
	m.Merge()
	fmt.Println(m)
}
