package structure

import "fmt"

type geometry interface {
	area() float64
	perim() float64
}

func measure(g geometry) {
	fmt.Println("Area = ", g.area())
	fmt.Println("Perimeter = ", g.perim())
}
