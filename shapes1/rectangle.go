package shapes1

import (
	"fmt"
	"log"
)

var Length = 2
var Width = 3

func Area(l, w int) int {
	return l * w
}
func init() {
	fmt.Println("=> init() from shapes1.rectangle package.")
	if (Length < 0) || (Width < 0) {
		log.Fatal("Length & Width cannot be less that zero")
	}
}
