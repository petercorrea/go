package pointers

import "fmt"

func Pointers() {
	fmt.Println("***Section 6: Pointers***")

	i := 100
	// to create a pointer, assign it the address of a variable
	p := &i

	// to reference the value of a pointer, use an asterisk
	fmt.Println(*p)
	*p = 200
	
	// the value has been updated
	fmt.Println(i)
}