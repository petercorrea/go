package loops

import "fmt"

// there is only one type of loop in Go...the For loop.
// no parenthesis needed
// the init and post statements are optional
func Loops() {
	fmt.Println("***Section 2: Loops***")
	fmt.Println("in control flow")

	sum := 0

	// for loop
	for i:=0; i<10;i++ {
		sum += i
	}

	fmt.Println(sum)

	sum = 1
	// while loop
	for sum < 1000 {
		sum += sum
	}

	// infinite loop
	// for {
	// 	fmt.Println("This will print forever")
	// }

	fmt.Println(sum)
}
