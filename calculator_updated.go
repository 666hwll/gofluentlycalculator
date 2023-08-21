package main

import (
	"fmt"
)

var operator string
var first_number float64
var second_number float64
var solution float64
var condition_main = true

func operations() {
	switch operator {
	case "+":
		solution = first_number + second_number

	case "-":
		solution = first_number - second_number

	case "*":
		solution = first_number * second_number

	case "/":
		if second_number == 0 {
			fmt.Println("Division through zero is not possible")
		} else {
			solution = first_number / second_number
		}

	default:
		fmt.Println("Invalid Input")

	}
}

func main() {
	for condition_main == true {
		fmt.Scan(&first_number)
		fmt.Scan(&operator)
		fmt.Scan(&second_number)
		operations()
		fmt.Println(*&solution)
	}
}
