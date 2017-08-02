package main

import (
	"net/http"
)

func main() {
	http.Handle("/", http.HandlerFunc(redirect))
	http.ListenAndServe(":3000", nil)
}

func redirect(w http.ResponseWriter, r *http.Request) {

	//TODO get path -> url map from DB/Redis
	switch r.URL.Path {
	case "/statistic":
		http.Redirect(w, r, "https://docs.google.com/forms/d/e/1FAIpQLSe3sMjvlpw90JlIwEDFGmZACnPDXRjA8KGDbKGY3xjBsvYRig/viewform", 307)
	case "/statistic-result":
		http.Redirect(w, r, "https://docs.google.com/spreadsheets/d/1gCxRH7Qs3Hoi9XIYKtJqwK32pw7xtFB6BBsSu7rjQ0c/edit#gid=1007446279", 307)
	}
}
