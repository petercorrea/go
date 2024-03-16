package functions

import "fmt"

// normal function
func Add(x int, y int) int {
	return x + y
}

// return multiple values
func Swap(x, y int) (int, int) {
	return y, x
}

// group paramters under a shared type
func Subtract(x, y int) int {
	return x - y
}

// named return values and naked returns
func Named(x, y int) (a, b int) {
	a = x + y
	b = x - y 
	return // naked return; automaticaly returns a and b
}

func Functions() {
	fmt.Println("***Section 5: Functions***")

	added := Add(1, 2)
	result:= fmt.Sprintf("Added result is: %d", added)
	fmt.Println(result)

	subtraced := Subtract(5, 3)
	result = fmt.Sprintf("Subtracted result is: %d", subtraced)
	fmt.Println(result)

	swapped_1, swapped_2 := Swap(1, 2)
	result = fmt.Sprintf("Swapped result is: %d %d", swapped_1, swapped_2)
	fmt.Println(result)

	named_1, named_2 := Named(5, 4)
	result = fmt.Sprintf("Nammed result is: %d %d", named_1, named_2)
	fmt.Println(result)
}
