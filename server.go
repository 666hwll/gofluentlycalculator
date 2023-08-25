package main

import (
	"fmt"
	"log"
	"net/http"
)

var first_numbers float64
var second_numbers float64
var operators string
var solutions float64

func operations_for_server() {

	switch operators {
	case "+":
		*&solutions = first_numbers + second_numbers

	case "-":
		*&solutions = first_numbers - second_numbers

	case "*":
		*&solutions = first_numbers * second_numbers

	case "/":
		if second_numbers == 0 {
			fmt.Println("Division through zero is not possible")
		} else {
			*&solutions = first_numbers / second_numbers
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
	first_numbers := r.FormValue("ValueOne")
	operators := r.FormValue("OperatorfCalc")
	second_numbers := r.FormValue("ValueTwo")

	fmt.Fprintf(w, "ValueOne = %s\n", first_numbers)
	fmt.Fprintf(w, "Operator = %s\n", operators)
	fmt.Fprintf(w, "ValueTwo = %s\n", second_numbers)
	//operations_for_server()
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
