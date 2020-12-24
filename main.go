package main

import (
	"html/template"
	"net/http"
)

var tpl = template.Must(template.ParseFiles("index.html", "result.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

func findPrimeHandler(w http.ResponseWriter, r *http.Request) {

	tpl.ExecuteTemplate(w, "result.html", 23)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/findPrime", findPrimeHandler)

	http.ListenAndServe(":8080", mux)
}
