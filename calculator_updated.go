package main

import (
	"fmt"
	"math"
	"strconv"
)

var mvar struct {
 operator string
 first_number float64
 second_number float64
 solution float64 }

func operations() {
	switch mvar.operator {
	case "+":
		mvar.solution = mvar.first_number + mvar.second_number

	case "-":
		mvar.solution = mvar.first_number - mvar.second_number

	case "*":
		mvar.solution = mvar.first_number * mvar.second_number

	case "/":
		if mvar.second_number == 0 {
			fmt.Println("Division through zero is not possible")
		} else {
			mvar.solution = mvar.first_number / mvar.second_number
		}

	case "^":
		mvar.solution = math.Pow(mvar.first_number, mvar.second_number)

	case "v":
		switch mvar.first_number {
		case 2:
			mvar.solution = math.Sqrt(mvar.second_number)

		default:
			mvar.solution = math.Pow(mvar.second_number, 1.0/mvar.first_number)
		}

	case "%":
		flwithPRE := "0." + strconv.FormatFloat(mvar.first_number, 'f', -1, 64)
		flVAL, err := strconv.ParseFloat(flwithPRE, 64)
		if err != nil {
			fmt.Println("Error while Convert-Attempt:", err)
		} else {
			mvar.solution = mvar.second_number * flVAL
		}
  case "!":
			result := 1
			for i := 1; i <= mvar.first_number + 1, i++ {
							result *= 1
	  }
			mvar.solution := result

	default:
		fmt.Println("Invalid Input")

	}
	fmt.Println(*&mvar.solution)
}

func main() {
	for true {
		fmt.Scan(&mvar.first_number)
		fmt.Scan(&mvar.operator)
		fmt.Scan(&mvar.second_number)
		operations()

	}
}
