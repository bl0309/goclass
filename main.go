package main

import (
	"fmt"
	//"error"
)

func main() {
	sumOfEvenOdd()
	maximumValue := max(8, 3)
	fmt.Println("max no: ", maximumValue)

	x, y := swap("Bikram", "Lamsal")
	fmt.Println("x and y value: ", x, y)

	a := 10
	b := 23

	fmt.Println("BS of a: ", a)
	fmt.Println("BS of b: ", b)

	swapAgain(&a, &b)

	fmt.Println("AS of a: ", a)
	fmt.Println("AS of b: ", b)

}

func sumOfEvenOdd() {
	sumEven := 0
	sumOdd := 0

	for i := 1; i < 50; i++ {
		if i%2 == 0 {
			sumEven += i
		} else {
			sumOdd += i
		}
	}
	fmt.Println(sumEven)
	fmt.Println(sumOdd)

}

func max(num1, num2 int) int {

	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}

func swap(x, y string) (string, string) {
	return y, x
}

func swapAgain(a, b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp
	fmt.Println("IS of a: ", a)
	fmt.Println("IS of b: ", b)
}
