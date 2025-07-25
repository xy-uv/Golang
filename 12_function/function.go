package main

import (
	"fmt"
)

// Normal Function
func add(a int, b int) int {
	sum := a + b
	return sum
}

func multi(a,b int)int{
	return a * b
}

// multiple return function

func getLang()(string,int,float32,bool,[]string){
	return "Golang",331,3.09,true,[]string{"I","am","learning","Golang"}
}

// Higher order function

func processSomething(fn func( a int)int)func(x int)int{
	return func(x int)int{
		result:= fn(x)

		return result
	}
}
func square(n int)int{
	return n*n
}

func main() {
	fmt.Println(add(11, 99))
	fmt.Println(multi(11, 11))
	_,_,_,_,array:=getLang()

	for _,ar :=range array{
		fmt.Println(ar)
	}

	squareFunc := processSomething(square)

	finalResult :=squareFunc(9)

	fmt.Println(finalResult)

	// inline function
	fn:= func(g int)int{
		return g
	}
	fmt.Println(fn(9))
	
	// anonyms function
	func(x,y,z float32){
		fmt.Println(x+y*z)
	}(3,7,2)
}