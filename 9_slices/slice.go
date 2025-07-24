package main

import (
	"fmt"
	"slices"
)

// slice -> a dynamically-sized, flexible view into the elements of an array.
// most used data structure/construct in Go.
// slices are more powerful than arrays, as they can grow and shrink in size.
// + most useful operations/methods on slices are built into the language (Go).
// + slices are reference types, meaning they point to an underlying array.
// + slices can be created using the built-in `make` function or by slicing an array.
// + slices have a length and a capacity. The length is the number of elements in the slice, and the capacity is the number of elements in the underlying array that can be accessed without reallocating memory.


func main() {
	// uninitialized slice is nil
	var ns []int
	fmt.Println(ns == nil)
	fmt.Println(len(ns))

	var nk = make([]int, 0, 5)
	// capacity -> maximum numbers of elements can fit
	fmt.Println(cap(nk))
	fmt.Println(nk == nil)

	nj := []int{}

	nj = append(nj, 1)
	nj = append(nj, 2)

	nj[0] = 3
	nj[1] = 5
	fmt.Println(nj)
	fmt.Println(cap(nj))
	fmt.Println(len(nj))

	var nas = make([]int, 0, 5)
	nas = append(nas, 2)
	var nas2 = make([]int, len(nas))

	// copy function
	copy(nas2, nas)
	fmt.Println(nas, nas2)

	// slice operator

	var nm = []int{1, 2, 3, 4, 5}
	fmt.Println(nm[0:1])
	fmt.Println(nm[:2])
	fmt.Println(nm[1:])

	// slices
	var nu1 = []int{1, 2, 3}
	var nu2 = []int{1, 2, 4}

	fmt.Println(slices.Equal(nu1, nu2))

	var nubs = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(nubs)

	numbers:=[]int{1,4,6,7,9,3,5}
	sum:=0
	for i :=range numbers{
		sum+=numbers[i]
	}
	fmt.Println(sum)

	names:=make([]string,0)

	names= append(names, "Doya","Samir","Ayat","Alif","Azan","Sitara")
	fmt.Println(names)
}