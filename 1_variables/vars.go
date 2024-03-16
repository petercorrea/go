package variables

import "fmt"

// ':=' is shorthand for variable declaration and initialization and is not allowed outside of functions
// variables are automatically initialized to their 'zero value', 0 for ints, false for booleans, "" for strings
// types can be inferred
// constants can never be declared with ':=', that's only for variables
// constant can be untyped and will define its type depending on the context it's used in
var age int
var isHuman bool
var name string
var hometown = "Miami"
const planet = "Earth"

func Variables() {
	fmt.Println("***Section 1: Variables***")
	fmt.Println(age, isHuman, name, hometown, planet)
}