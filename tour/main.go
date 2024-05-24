package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func main() {
	//For is While
	sum := 1
	for sum < 1000 {
		sum += sum
	}

	//Go's if statements are like its for loops; the expression need not be surrounded by parentheses ( ) but the braces { } are required.

	x := 1
	if x < 0 {
		fmt.Println(x)
	}
	pow(2, 4, 8)
	testSwitch()
	testSwitch2()
	testDefer()

}

func pow(x, n, lim float64) float64 {
	//Variables declared by the statement are only in scope until the end of the if.
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		//Variables declared inside an if short statement are also available inside any of the else blocks.
		fmt.Printf("%g <= %g\n", v, lim)
	}
	return lim
}

func testSwitch() {
	//the break statement that is needed at the end of each case in those languages is provided automatically in Go.
	//Another important difference is that Go's switch cases need not be constants, and the values involved need not be integers.
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

func testSwitch2() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func testDefer() {
	defer fmt.Println("world")

	fmt.Println("hello")
}
