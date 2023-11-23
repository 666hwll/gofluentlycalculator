package main

import (
	"fmt"
	"math"
)

var mvar struct {
	oper string
	fnum float64
	snum float64
	solu float64
	svst float64
}

func opera(x float64, y string, z float64, a float64, b float64) float64 {
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
			print("\x1b[31mdivision through 0 is impossible\x1b[0m\n")
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

	case "round":
		switch z {
		case 0:
			a = math.Round(b)
		default:
			a = math.Round(x)
		}

	case "help":
		print("\x1b[31mFormat: Number Operator Number; more in doc.txt\x1b[0m\n")

	case "exit":
		return 0

	default:
		print("\x1b[31mInvalid input. Type help\x1b[0m\n")
	}
	b = a
	return a
}

func main() {
	for {
		fmt.Scan(&mvar.fnum, &mvar.oper, &mvar.snum)
		fmt.Println(opera(mvar.fnum, mvar.oper, mvar.snum, mvar.solu, mvar.svst))
	}

}
