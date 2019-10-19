package shapes2

import "fmt"

func Area(l, w int) int {
	return l * w * 2
}
func init() {
	fmt.Println("=> init() from shapes2.rectangle package")
}
