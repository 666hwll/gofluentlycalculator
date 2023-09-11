package main

import (
	"fmt"
	"math"
	"strconv"
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

	case "^":
		solution = math.Pow(first_number, second_number)

	case "v":
		switch first_number {
		case 2:
			solution = math.Sqrt(second_number)

		default:
			solution = math.Pow(second_number, 1.0/first_number)
		}

	case "%":
		flwithPRE := "0." + strconv.FormatFloat(first_number, 'f', -1, 64)
		flVAL, err := strconv.ParseFloat(flwithPRE, 64)
		if err != nil {
			fmt.Println("Error while Convert-Attempt:", err)
		} else {
			solution = second_number * flVAL
		}

	default:
		fmt.Println("Invalid Input")

	}
	fmt.Println(*&solution)
}

func main() {
	for condition_main == true {
		fmt.Scan(&first_number)
		fmt.Scan(&operator)
		fmt.Scan(&second_number)
		operations()

	}
}
