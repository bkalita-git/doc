package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Length float64
	Width  float64
}

func (circle Circle) Area() float64 {
	return 2 * math.Pi * circle.Radius
}
func (rectangle Rectangle) Area() float64 {
	return rectangle.Length * rectangle.Width
}

type AreaInterface interface {
	Area() float64
}

func generic[V int | string](var1 V) V { //V for variable
	print("\nhit")
	var s V
	return s
}

func hook(view interface{}) {
	switch view.(type) {
	case func(string) string:
		fmt.Printf("\n1 param")

		// case func(string, string) string:
		// 	fmt.Printf("\n2 param")
	}
	fmt.Printf("\n%T", view)
}

func view1(req string) string {
	return "view1"
}

func view2(req string, var1 string) string {
	return "view2"
}

func main() {
	return
	fmt.Printf("\n%v", generic("ok"))
	hook(view1)
	hook(view2)
	return
	c1 := Circle{Radius: 30}
	r1 := Rectangle{Length: 30, Width: 30}
	// fmt.Printf("\nArea of circle %v", c1.Area())
	// fmt.Printf("\nArea of rectangle %v", r1.Area())
	var a AreaInterface = c1
	fmt.Printf("\nArea of circle %v", a.Area())

	var b AreaInterface = r1

	fmt.Printf("\nArea of rectangle %v", b.Area())

}
