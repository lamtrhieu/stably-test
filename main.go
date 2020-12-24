package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

var tpl = template.Must(template.ParseFiles("index.html", "result.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

func findPrimeHandler(w http.ResponseWriter, r *http.Request) {
	input := r.FormValue("input")
	fmt.Println(input)

	i, error := strconv.Atoi(input)

	if error != nil {
		tpl.ExecuteTemplate(w, "index.html", error)
	}

	p := FindPrime(i)

	tpl.ExecuteTemplate(w, "result.html", p)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/findPrime", findPrimeHandler)

	http.ListenAndServe(":8080", mux)
}
