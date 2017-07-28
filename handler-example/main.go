package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	rh := http.RedirectHandler("https://docs.google.com/spreadsheets/d/1gCxRH7Qs3Hoi9XIYKtJqwK32pw7xtFB6BBsSu7rjQ0c/edit?usp=sharing", 307)

	mux.Handle("/statistic", rh)

	log.Println("Listening...")

	http.ListenAndServe(":3000", mux)
}
