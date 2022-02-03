package goroutines

import (
	"errors"
	"fmt"
)

func InitErrorStrategy() {
	value, err := f1(5)
	if err != nil {
		panic(err)
	}
	fmt.Println("value = ", value)
	// n, err := fmt.Println("value = ", value)
	// fmt.Println(n, err)

}

func f1(arg int) (int, error) {
	if arg < 0 {
		return 0, errors.New("Can't work with negative numnbers")
	}
	return arg + 1, nil
}
