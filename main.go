package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this into a new Article struct
	// append this to our Articles array.
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)
	// update our global Articles array to include
	// our new Article
	Articles = append(Articles, article)

	json.NewEncoder(w).Encode(article)
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
