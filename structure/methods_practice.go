package structure

import (
	"fmt"
	"math"
)

type rect struct {
	width  float64
	height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.height * r.width
}

func (r rect) perim() float64 {
	return 2 * r.height * 2 * r.width
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func Init2() {
	r := rect{width: 10, height: 3}
	c := circle{radius: 5}
	// r2 := rect{width: 65, height: 78}
	// fmt.Println("Area = ", r.height*r.width)

	fmt.Println("Area of rectangle = ", r.area())
	fmt.Println("Perimeter of rectangle = ", r.perim())

	// r2 := &r
	// fmt.Println("Area of rectangle = ", r2.area())
	// fmt.Println("Perimeter of rectangle = ", r2.perim())

	// call function for circle
	fmt.Println("Area of circle = ", c.area())
	fmt.Println("Perimeter of cirlce = ", c.perim())

	// interface call
	fmt.Println("--------Rectangle----------")
	measure(r)
	fmt.Println("--------Circle-------------")
	measure(c)
}
