package main

import (
	"fmt"
	"time"
)

func main() {
	// switch is a control structure that allows you to execute different code blocks based on the value of an expression.
	// It is similar to the switch statement in other programming languages.

	// Basic switch statement
	switch day := 3; day {
	case 1:
		println("Monday")
	case 2:
		println("Tuesday")
	case 3:
		println("Wednesday")
	case 4:
		println("Thursday")
	case 5:
		println("Friday")
	default:
		println("Weekend")
	}

	// Switch with multiple cases
	switch month := "March"; month {
	case "January", "February", "March":
		println("First quarter of the year")
	case "April", "May", "June":
		println("Second quarter of the year")
	default:
		println("Other months")
	}

	// Switch with no condition (acts like if-else)
	switch {
	case true:
		println("This is always true")
	case false:
		println("This will never be printed")
	default:
		println("Default case")
	}

	// switch with multiple conditions
	switch time.Now().Weekday(){
		case time.Saturday, time.Sunday:
			println("It's the weekend!")		
		case time.Monday, time.Tuesday, time.Wednesday, time.Thursday:
			println("It's a workday.")
		case time.Friday:
			println("It's almost the weekend/Praying day!")
	}
		
	// type switch
	whatType:=func(i any){
		switch i.(type){
		case int:
			fmt.Println("It's an integer")
		case string:
			fmt.Println("It's a string")
		case bool:
			fmt.Println("It's a boolean")
		case float32,float64:
			fmt.Println("It's a float")
		default:
			fmt.Println("Other")
	}
}
whatType("sanda")
}