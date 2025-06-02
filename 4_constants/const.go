package main

import "fmt"

// Declare variable outside the function

const age=33
var name="Golang"

// age:=33 not possible

func main(){
	const nab="Nab"
//  fmt.Println(age)


// Const grouping
	const (
		port=5000
		host="localhost"
	)
	fmt.Println(host,port)
}