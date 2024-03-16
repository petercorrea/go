package control_flow

import (
	"fmt"
	"runtime"
	"time"
)

func ControlFlow() {
	fmt.Println("***Section 3: Control Flow***")

	x := 1
	// defer runs after the containing function completes
	// defers are placed onto a stack and executed in last-in-first-out order
	// arguments are evaluated immediantly
	msg := "world"
	defer fmt.Println(msg)
	msg = "hello"
	defer fmt.Println(msg)
	msg = "hi" // never printed

	if x < 10 {
		fmt.Println("x is smaller than 10")
	} else if x > 10 && x < 20 {
		fmt.Println("x is smaller than 20")
	} else {
		fmt.Println("x is greater than 20")
	}

	// if statements may begin with a short statement before the conditional
	if y:= 2; y < 10 {
		fmt.Println("y is smaller than 10")
	}

	// switch statements automatically provide fall-through protection
	// similar to if statements, may begin with a statement
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd, plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	// switch statements can omit the conditional
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