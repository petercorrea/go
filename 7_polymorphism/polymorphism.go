// structs, pointer receivers, interfaces, generics

package polymorphism

import "fmt"

// a struct is a set of fields
type Vertex struct {
	x int
	y int
}

// a method is a function with a receiver argument
	// the method and the receiver must be defined within the same package
	// the receiver is passed by value
func (v Vertex) printX() {
	fmt.Println(v.x)
}

// to pass the receiver by reference, use pointer receivers
	// use pointer receivers to modify or to reduce copying large objects
	// as a convenience, Go will intepret (v).addTwo() == (&v).addTwo()
func (v *Vertex) addTwo() {
	v.x += 3
}

// an interface is a set of method signatures
	// interfaces are implemented implicitly
type Ops interface {
	printX()
	AddTwo() int
}

// for generic programming we have generic functions and generic types
// type parameters come before parameters
// 'any' is a type contraint, there are others
func swap[T any](a, b T) (T, T) {
	return b, a
}

func Structs() {
	fmt.Println("***Section 7: Polymorphism***")

	my_struct := Vertex{1, 2}

	// calling a method with a value receiver
	my_struct.printX()

	// calling a method with a pointer receiver
	my_struct.addTwo()

	fmt.Println(my_struct.x)

	x, y := swap(1, 2)
	a, b := swap("Peter", "Correa")

	fmt.Println(x, y)
	fmt.Println(a, b)
}