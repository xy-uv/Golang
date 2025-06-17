package main

import "fmt"

func main() {
	age:=61

	if age < 18 {
		fmt.Println("You are a minor.")
	} else if age >= 18 && age < 60 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a senior citizen.")
	}

	if mark:=69; mark>=80 {
		fmt.Println("You got an A grade.")
	} else if mark >= 60 {
		fmt.Println("You got a B grade.")
	} else if mark >= 40 {
		fmt.Println("You got a C grade.")
	} else {
		fmt.Println("You failed.")
	}
}