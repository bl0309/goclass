package goroutines

import (
	"fmt"
	"time"
)

func Init() {
	// routineMain()

	// channels

	mChannal := make(chan string)
	mChannal <- "mainThread"

	mChannal <- "Another thread"
	go func(msg chan string) {
		mChannal <- "goroutine"
	}(mChannal)

	go routine1(mChannal)
	// go routine2(mChannal)

	msg := <-mChannal
	fmt.Println(msg)
	msg = <-mChannal
	fmt.Println(msg)
	msg = <-mChannal
	fmt.Println(msg)
}

func routine1(msg chan string) {
	fmt.Println(<-msg)
	msg <- "routine1"
}

// func routine2(msg <-chan string) {
// 	fmt.Println(<-msg)
// 	// msg <- "routine2"
// }

func routineMain() {
	fmt.Println("First")
	routine()
	fmt.Println("second")
	go routine()

	go func() {
		fmt.Println("Third")
	}()

	go func(msg string) {
		fmt.Println(msg)
	}("testing")

	time.Sleep(time.Second)
	fmt.Println("done")

}

func routine() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}
