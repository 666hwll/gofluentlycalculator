package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var oval struct {
	first_numbers  float64
	second_numbers float64
	operators      string
	solutions      float64
}

func operations_for_server(w http.ResponseWriter) {

	switch oval.operators {
	case "+":
		oval.solutions = oval.first_numbers + oval.second_numbers

	case "-":
		oval.solutions = oval.first_numbers - oval.second_numbers

	case "*":
		oval.solutions = oval.first_numbers * oval.second_numbers

	case "/":
		if oval.second_numbers == 0 {
			fmt.Println("Division through zero is not possible")
		} else {
			oval.solutions = oval.first_numbers / oval.second_numbers
		}

	default:
		fmt.Fprintf(w, "Invalid Input\n")

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
	var err error
	if err = r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	// The submit button was pressed
	fmt.Fprintf(w, "POST request successful!\n")
	firstnumbers := r.FormValue("ValueOne")
	oval.operators = r.FormValue("OperatorfCalc")
	secondnumbers := r.FormValue("ValueTwo")
	//submitValue := r.FormValue("submit")
	//if submitValue != "" {
	oval.first_numbers, err = strconv.ParseFloat(firstnumbers, 64)
	if err != nil {
		fmt.Fprintf(w, "error while converting first float\n")
	}
	oval.second_numbers, err = strconv.ParseFloat(secondnumbers, 64)
	if err != nil {
		fmt.Fprintf(w, "error while converting second float")
	}
	operations_for_server(w)
	fmt.Fprintf(w, "%f", oval.solutions)
	//}

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
