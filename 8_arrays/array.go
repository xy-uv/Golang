package main

import "fmt"

// numbered sequence of specific length
func main() {
	var numbers [5]int // Declare an array of integers with a fixed size of 5
	numbers[0] = 10
	numbers[1] = 20

	fmt.Println("Array of numbers:", numbers)

	var floats [3]float64 // Declare an array of floats with a fixed size of 3
	floats[0] = 1.1

	fmt.Println("Array of floats:", floats)

	var fruits [3]string // Declare an array of strings with a fixed size of 3
	fruits[1] = "Banana"

	fmt.Println("Array of fruits:", fruits)

	var bools [2]bool // Declare an array of booleans with a fixed size of 2
	bools[0] = true

	fmt.Println("Array of booleans:", bools)

	var mixed [7]any // Declare an array of empty interface to hold mixed types
	mixed[0] = 42
	mixed[2] = "Hello"
	mixed[4] = 3.14
	mixed[6] = false

	fmt.Println("Array of mixed types:", mixed)

	// Declare and assign values to an array in one line
	colors := [4]string{ "Green", "Blue", }
	mixedColors := [3]any{ "Red", 255, true }
	nums:= [5]int{3,5,7}
	fmt.Println("Array of colors:", colors)
	fmt.Println("Array of mixed colors:", mixedColors)
	fmt.Println("Array of numbers:", nums)

	// Accessing array length with len() function
	fmt.Println("Length of numbers array:", len(numbers))
	fmt.Println("Length of floats array:", len(floats))
	fmt.Println("Length of fruits array:", len(fruits))
	fmt.Println("Length of bools array:", len(bools))
	fmt.Println("Length of mixed array:", len(mixed))
	fmt.Println("Length of colors array:", len(colors))
	fmt.Println("Length of mixedColors array:", len(mixedColors))
	fmt.Println("Length of nums array:", len(nums))


	// 2D arrays
	matrix := [2][2]int{
		{1, 3},
		{5, 1},
	}
	fmt.Println("2D Array (Matrix):", matrix)

	// Array use case in go
	// // 1. Fixed size: Arrays have a fixed size, which means you need to know the number of elements in advance.
	// // 2. Homogeneous: All elements in an array must be of the same type.
	// // 3. Memory efficiency: Arrays are more memory-efficient than slices because they have
	// a fixed size and do not require additional memory allocation.
	// // 4. Performance: Arrays can be more performant than slices in certain scenarios,
	// especially when the size is known at compile time.
	// // // 5. Use cases: Arrays are often used in scenarios where the size of the collection is known
	// // at compile time, such as in mathematical computations, image processing, or when working
	// // with fixed-size data structures.
	// // 6. Syntax: Arrays are defined using the [size]Type syntax,
	// //7. Constant time access: Arrays provide constant time access to elements,
	// // which means you can access any element in the array in O(1) time
	

}