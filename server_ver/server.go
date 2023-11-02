package main

import (
	"fmt"
	"log"
	"net/http"
)

var first_numbers float64
var	second_numbers float64
var	operators string
var	solutions float64

func operations_for_server() {

	switch operators {
	case "+":
		solutions = first_numbers + second_numbers

	case "-":
		solutions = first_numbers - second_numbers

	case "*":
		solutions = first_numbers * second_numbers

	case "/":
		if second_numbers == 0 {
			fmt.Println("Division through zero is not possible")
		} else {
			solutions = first_numbers / second_numbers
		}

	default:
		fmt.Println("Invalid Input")

	}
}

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/Welcome" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello World!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful!\n")
	firstnumbers := r.FormValue("ValueOne")
	operator := r.FormValue("OperatorfCalc")
	secondnumbers := r.FormValue("ValueTwo")
	&first_numbers, &operators, &second_numbers := strconv.ParseFloat(firstnumbers, operator, secondnumbers, 64)
	fmt.Fprintf(w, "First number = %s\n", firstnumbers)
	fmt.Fprintf(w, "Operator = %s\n", operator)
	fmt.Fprintf(w, "Second number = %s\n", secondnumbers)
	operations_for_server()
	//fmt.Fprintf(w, "solution: %f\n", *&solutions)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/Welcome", WelcomeHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
