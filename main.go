package main

import (
	"fmt"

	p1 "github.com/pseudonative/goStuffAll/animals"

	"github.com/pseudonative/goStuffAll/shapes1"
	"github.com/pseudonative/goStuffAll/shapes2"
)

func main() {
	testCase1()
	testCase2()
	// testCase3()

}
func testCase1() {
	fmt.Printf("shape1 pkg - Rectangle Area= %d\n", shapes1.Area(shapes1.Length, shapes1.Width))
	fmt.Printf("shape1 pkg - Square Area= %d\n", shapes1.SquareArea(2))
	fmt.Printf("shape2 pkg - Rectangle Area= %d\n", shapes2.Area(2, 3))
	fmt.Printf("shape2 pkg - Square Area= %d\n", shapes2.SquareArea(2))

}

func testCase2() {
	lion := p1.Lion{}
	fmt.Println(lion.Speaks())

	cat := p1.Cat{}
	fmt.Println(cat.Speaks())

	human := p1.Human{}
	fmt.Println(human.Speaks())
}

// func testCase3() {

// }
// func changeAthleteName1(p athletes.Player) {

// }
// func changeAthleteName2(p *athletes.Player) {

// }
