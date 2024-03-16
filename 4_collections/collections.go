package collections

import "fmt"

func Arrays() {
	fmt.Println("***Section 4: Collections***")
	// arrays are of type [n]T
	// array are fixed size
	my_array := [4]int{1, 2, 3, 4}
	fmt.Println(my_array)

	// slices are of type []T
	// slice are dynamically sized references to arrays
	// the bounds of a slic are optional
	my_slice := my_array[0:2]
	fmt.Println(my_slice)

	// you can initialize arrays by using slices and omiting the length
	// this creates an array and then a slice that references it
	other_array := []int{100, 200, 300}
	fmt.Println(other_array)

	// slices have a length and a capacity, 
	// where the length is the length of the slice
	// and capacity is the length of the array it points to
	// you can resize a slice by reslicing,
	// but you must stay eith the capacity of the array
	// outOfBounds err

	// the zero value of a slice is nil, with 0 capacity and length 
	// and no reference to an underlying array

	// one can create a slice using the make function
	// make([]T, len, cap)

	// we can append to a slice by using append()
	// append([]T, ...T)

	// ranges return idx, ele

	// maps
	// map[K]T
	// make(map[K]T)
	my_map := map[string]int{"a": 1, "b":2}
	fmt.Println(my_map)

	// delete elemetn using delete(m, k)
	// check fo presence using ele, ok := m[k]
}