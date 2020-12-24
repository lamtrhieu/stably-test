package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Will show something here.")
}

func findPrimeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Will find prime here")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/findPrime", findPrimeHandler)

	http.ListenAndServe(":8080", mux)
}
