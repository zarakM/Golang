package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func add(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	one := vars["one"]
	two := vars["two"]

	var (
		first_operand  int
		second_operand int
	)
	if i, err := strconv.Atoi(one); err == nil {
		first_operand = i
	}

	if i, err := strconv.Atoi(two); err == nil {
		second_operand = i
	}

	result := strconv.Itoa(first_operand + second_operand)
	fmt.Fprintf(w, "Result is %v", result)
}

func subtract(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	one := vars["one"]
	two := vars["two"]

	var (
		first_operand  int
		second_operand int
	)
	if i, err := strconv.Atoi(one); err == nil {
		first_operand = i
	}

	if i, err := strconv.Atoi(two); err == nil {
		second_operand = i
	}

	result := strconv.Itoa(first_operand - second_operand)
	fmt.Fprintf(w, "Result is %v", result)
}

func multiply(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	one := vars["one"]
	two := vars["two"]

	var (
		first_operand  int
		second_operand int
	)
	if i, err := strconv.Atoi(one); err == nil {
		first_operand = i
	}

	if i, err := strconv.Atoi(two); err == nil {
		second_operand = i
	}

	result := strconv.Itoa(first_operand * second_operand)
	fmt.Fprintf(w, "Result is %v", result)
}

func divide(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	one := vars["one"]
	two := vars["two"]

	var (
		first_operand  int
		second_operand int
	)
	if i, err := strconv.Atoi(one); err == nil {
		first_operand = i
	}

	if i, err := strconv.Atoi(two); err == nil {
		second_operand = i
	}

	result := strconv.Itoa(first_operand / second_operand)
	fmt.Fprintf(w, "Result is %v", result)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/add/{one}/{two}", add)
	myRouter.HandleFunc("/subtract/{one}/{two}", subtract)
	myRouter.HandleFunc("/multiply/{one}/{two}", multiply)
	myRouter.HandleFunc("/divide/{one}/{two}", divide)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	handleRequests()
}
