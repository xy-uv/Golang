package main

import "fmt"

func main() {
	// explicit type
	var nam string = "Golang"
	fmt.Println(nam)

	//infer type
	var name="Golang"
	fmt.Println(name)

	var num=33.0/2
	fmt.Println(num)

	var age =33/2
	fmt.Println(age)

	// Variable grouping
	var (
		nameHer="Senaha"
		name_her="Sanha"
	)
	fmt.Println(nameHer,name_her)
	// Shorthand variables define with assign
	nan:=37./2
	fmt.Println(nan)
}