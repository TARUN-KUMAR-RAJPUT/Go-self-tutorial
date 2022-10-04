package main

import (
	"fmt"
	"net/http"
)


func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Yoo Man")
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is Home")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is about")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/home", home)
	http.HandleFunc("/about", about)

    http.ListenAndServe(":8080", nil)
}