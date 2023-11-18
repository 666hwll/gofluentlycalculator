package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"strconv"
)

var mvar struct {
	oper string
	fnum float64
	snum float64
	solu float64
	svst float64
}

var PrOPT struct {
	prtstandpr uint
	stdpr      uint
}

type Proset struct {
	Precision uint `json:"precision"`
}

func openfl() int {
	hmdir, err := os.UserHomeDir()
	if err != nil {
		return 3
	}
	dirp := filepath.Join(hmdir, "/.gocalc/settings.json")
	jsonFile, err := os.Open(dirp)
	if err != nil {
		fmt.Println(err)
		PrOPT.prtstandpr = PrOPT.stdpr
		return 1
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var option Proset
	err = json.Unmarshal(byteValue, &option)
	if err != nil {
		fmt.Println(err)
		PrOPT.prtstandpr = PrOPT.stdpr
		return 2
	}
	PrOPT.prtstandpr = option.Precision
	return 0
}

func rndFl(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func opera(x float64, y string, z float64, a float64, b float64) string {
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

	case "round":
		switch x {
		case 0:
			a = rndFl(b, PrOPT.prtstandpr)
		default:
			a = rndFl(x, uint(z))
		}

	case "help":
		return "Format: Number Operator Number; more in doc.txt"

	case "exit":
		return "0"

	default:
		return "Invalid input. Type help"
	}
	b = a
	return strconv.FormatFloat(a, 'f', int(PrOPT.prtstandpr), 64)
}

func main() {
	PrOPT.stdpr = 5
	openfl()
	flag.StringVar(&mvar.oper, "o", "", "operation")
	flag.Float64Var(&mvar.fnum, "f", 0, "firstnum")
	flag.Float64Var(&mvar.snum, "s", 1, "secondnum")
	flag.Parse()
	if mvar.oper != "" && mvar.fnum != 0 {
		fmt.Println(opera(mvar.fnum, mvar.oper, mvar.snum, mvar.solu, mvar.svst))
	} else {
		for {
			fmt.Scan(&mvar.fnum, &mvar.oper, &mvar.snum)
			fmt.Println(opera(mvar.fnum, mvar.oper, mvar.snum, mvar.solu, mvar.svst))
		}
	}

}
