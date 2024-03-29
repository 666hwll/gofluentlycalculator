package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
)

var mvar struct { //values used for processing
	oper string  // operator
	fnum float64 // first number of current operation
	snum float64 // second number of curr operation
	solu float64 // solution
	svst float64 // savestate
}

var PrOPT struct { //standart values
	rt         bool   // runtime
	prtstandpr uint   // standart precision point
	stdpr      uint   // default standart precision point
	col        string // text color
	msg        string // message(s)
	ct         uint   // counter
}

type Proset struct { //settings.json template
	Precision uint   `json:"precision"`
	Color     string `json:"col"`
}

func openfl() int { //opens, reads & parses the settings.json file
	hmdir, err := os.UserHomeDir()
	if err != nil {
		return 3
	}

	dirp := filepath.Join(hmdir, ".config", "gocalc", "settings.json")

	byteValue, err := os.ReadFile(dirp)
	if err != nil {
		fmt.Println(err)
		PrOPT.prtstandpr = PrOPT.stdpr
		return 1
	}

	var option Proset
	err = json.Unmarshal(byteValue, &option)
	if err != nil {
		fmt.Println(err)
		return 2
	}
	PrOPT.prtstandpr = option.Precision
	PrOPT.col = option.Color
	return 0
}

func preprocc() string {
	mvar.solu = opera(mvar.fnum, mvar.oper, mvar.snum, mvar.solu)
	//working on a new preproccessor to round up the UX of the programm a bit and the readability
	if PrOPT.msg == "" {
		PrOPT.msg = strconv.FormatFloat(mvar.solu, 'f', int(PrOPT.prtstandpr), 64)
		return PrOPT.col + PrOPT.msg + "\x1b[0m"
	}
	return ""

}

func rndFl(val float64, precision uint) float64 { //round function for opera()
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func inviformat() string { //error function for opera()
	var imsg string = "Invalid input. Type '0 help 0'"
	PrOPT.ct++
	if PrOPT.ct%2 == 0 {
		imsg = "'0 list 0' for commands and '0 exit 0' to exit."
	}
	return imsg
}

func opera(x float64, y string, z float64, a float64) float64 { // calculates, saves the outcome and convert to string
	switch y { // will later on be replaced with a dictionary and functions for each cmd
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
			PrOPT.msg = "\x1b[31mdivision through 0 is impossible\x1b[0m"
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
		a = math.Tan(x)
		z = float64(PrOPT.prtstandpr)

	case "atan":
		a = math.Atan(x)
		z = float64(PrOPT.prtstandpr)

	case "tanh":
		a = math.Tanh(x)
		z = float64(PrOPT.prtstandpr)

	case "atanh":
		a = math.Atanh(x)
		z = float64(PrOPT.prtstandpr)

	case "sin":
		a = math.Sin(x)
		z = float64(PrOPT.prtstandpr)

	case "asin":
		a = math.Asin(x)
		z = float64(PrOPT.prtstandpr)

	case "sinh":
		a = math.Sinh(x)
		z = float64(PrOPT.prtstandpr)

	case "asinh":
		a = math.Asinh(x)
		z = float64(PrOPT.prtstandpr)

	case "cos":
		a = math.Cos(x)
		z = float64(PrOPT.prtstandpr)

	case "acos":
		a = math.Acos(x)
		z = float64(PrOPT.prtstandpr)

	case "cosh":
		a = math.Cosh(x)
		z = float64(PrOPT.prtstandpr)

	case "acosh":
		a = math.Acosh(x)
		z = float64(PrOPT.prtstandpr)

	case "log":
		a = math.Log(x) / math.Log(z)

	case "%":
		a = z / 100 * x

	case "!":
		a = math.Gamma(x+1) * z

	case "round":
		switch x {
		case 0:
			a = rndFl(mvar.svst, PrOPT.prtstandpr)
		default:
			a = rndFl(x, uint(z))
		}

	case "help":
		PrOPT.msg = "Format: [Number][Space][Operator][Space][Number]"

	case "list":
		PrOPT.msg = "+ - * and x /\n^ v\n a/tan/h a/sin/h a/cos/h\n log % !\nround help exit"

	case "exit":
		PrOPT.rt = false

	default:
		PrOPT.msg = "\x1b[31m" + inviformat() + "\x1b[0m"
	}
	mvar.svst = a
	if a == 0 {
		fmt.Println(PrOPT.msg)
	}
	return a
}

func main() {
	PrOPT.rt = true
	PrOPT.stdpr = 5
	PrOPT.ct = 0
	flag.StringVar(&mvar.oper, "o", "", "operation")
	flag.Float64Var(&mvar.fnum, "f", 0, "firstnum")
	flag.Float64Var(&mvar.snum, "s", 1, "secondnum")
	flag.Parse()
	if mvar.oper != "" && mvar.fnum != 0 {
		fmt.Println(opera(mvar.fnum, mvar.oper, mvar.snum, mvar.solu))
	} else {
		for PrOPT.rt == true {
			mvar.solu = 0
			PrOPT.msg = ""
			go openfl()
			fmt.Scan(&mvar.fnum, &mvar.oper, &mvar.snum)
			//fmt.Println(opera(mvar.fnum, mvar.oper, mvar.snum, mvar.solu))
			fmt.Println(preprocc())

		}

	}
}
