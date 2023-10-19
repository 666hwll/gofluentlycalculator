package main

import (
	"fmt"
	"math"
)

var mvar struct {
 oper string
 fnum float64
 snum float64
 solu float64 }

func factorial(n float64, b float64) float64 {
	return math.Gamma(n+1) * b
}

func opera() {
	switch mvar.oper {
	case "+":
		mvar.solu = mvar.fnum + mvar.snum

	case "-":
		mvar.solu = mvar.fnum - mvar.snum

	case "*":
		mvar.solu = mvar.fnum * mvar.snum

	case "/":
		if mvar.snum == 0 {
			fmt.Println("Division through zero is not possible")
		} else {
			mvar.solu = mvar.fnum / mvar.snum
		}

	case "^":
		mvar.solu = math.Pow(mvar.fnum, mvar.snum)

	case "v":
		switch mvar.fnum {
		case 2:
			mvar.solu = math.Sqrt(mvar.snum)

		default:
			mvar.solu = math.Pow(mvar.snum, 1.0/mvar.fnum)
		}

	case "%":
		mvar.solu = mvar.snum / 100 * mvar.fnum
		
  	case "!":
		mvar.solu = factorial(mvar.fnum, mvar.snum)

	default:
		print("Invalid Input\n")

	}
	fmt.Println(*&mvar.solu)
}

func main() {
	for true {
		fmt.Scan(&mvar.fnum, &mvar.oper, &mvar.snum)
		opera()

	}
}
