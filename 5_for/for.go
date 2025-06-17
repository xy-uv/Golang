package main

import "fmt"

// for-> for is the only one constructor for lopping in go

func main(){
// while fashion type loop
	i:=1
	for (i<=7){
		fmt.Println(i)
		i+=2
	}

// infinite loop
	// for{
		fmt.Println("inf")
	// }

// classic for loop
	for j:=0;j<=9;j++{
		fmt.Println(j)
	}

// continue/break with loop
	for k:=0;k<=99;k+=2{
		if (k==44){
			// continue
			break
		}
		fmt.Println(k)
	}

// go 1.22 updated for loop with range
	for m:=range 100{
		fmt.Println(m)
	}
}