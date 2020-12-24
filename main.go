package main

import (
	"html/template"
	"net/http"
	"strconv"
)

var tpl = template.Must(template.ParseFiles("index.html", "result.html"))

var cache = make(map[int]int)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

func saveCache(start int, p int) {
	for i := start; i >= p; i-- {
		cache[i] = p
	}
}

func findPrimeHandler(w http.ResponseWriter, r *http.Request) {
	input := r.FormValue("input")

	i, error := strconv.Atoi(input)

	if error != nil || i < 2 {
		e := true
		tpl.ExecuteTemplate(w, "index.html", e)

	} else {
		p := cache[i]
		if p == 0 {
			p = FindPrime(i)

			saveCache(i, p)
		}

		tpl.ExecuteTemplate(w, "result.html", p)
	}

}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/findPrime", findPrimeHandler)

	http.ListenAndServe(":8080", mux)
}
