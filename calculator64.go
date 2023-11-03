package main

import (
	"fmt"
	"math"
	"flag"
	"strconv"
)

var mvar struct {
	oper string
	fnum float64
	snum float64
	solu float64
}
	
func opera(x float64, y string, z float64, a float64) string {
	switch y {
	case "+":
		a = x + z

	case "-":
		a = x - z

	case "*":
		a = x * z
	
	case "x":
		a = x * z

	case "/":
		if z == 0 {
			return "division through 0 is impossible"
		} else {
			a = x / z
		}

	case "^":
		a = math.Pow(x, z)

	case "v":
		switch x {
		case 2:
			a = math.Sqrt(z)
	
		default:
			a = math.Pow(z, 1.0/x)
		}
	
	case "tan":
		a = math.Tan(x) * z

	case "tanh":
		a = math.Tanh(x) * z
		
	case "sin":
		a = math.Sin(x) * z

	case "sinh":
		a = math.Sinh(x) * z

	case "cos":
		a = math.Cos(x) * z

	case "cosh":
		a = math.Cosh(x) * z

	case "log":
		a = math.Log(x) / math.Log(z)

	case "%":
		a = z / 100 * x

	case "!":
		a = math.Gamma(x+1) * z

	case "help":
		return "Format: Number Operator Number; more in doc.txt"

	default:
		return "Invalid input. Type help"
	}
	return strconv.FormatFloat(a, 'f', 7, 64)
}

func main() {
	flag.StringVar(&mvar.oper, "o", "", "operation")
	flag.Float64Var(&mvar.fnum, "f", 0, "firstnum")
	flag.Float64Var(&mvar.snum, "s", 1, "secondnum")
	flag.Parse()
	if (mvar.oper != "" && mvar.fnum != 0) {
		fmt.Println(opera(mvar.fnum, mvar.oper, mvar.snum, mvar.solu))
	} else {
		for {
			fmt.Scan(&mvar.fnum, &mvar.oper, &mvar.snum)
			fmt.Println(opera(mvar.fnum, mvar.oper, mvar.snum, mvar.solu))
		}
	}
	
}
