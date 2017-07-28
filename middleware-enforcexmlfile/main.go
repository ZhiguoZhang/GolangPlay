package main

import (
	"log"
	"bytes"
	"net/http"
)

func final(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("OK, It is good XML format\n"))
}

func enforceXMLHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ContentLength == 0 {
			http.Error(w, http.StatusText(400), 400)
			return
		}

		buf := new (bytes.Buffer)
		buf.ReadFrom(r.Body)
		if http.DetectContentType(buf.Bytes()) != "text/xml; charset=utf-8" {
			http.Error(w, http.StatusText(415), 415)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main()  {
	finalHandler := http.HandlerFunc(final)

	http.Handle("/", enforceXMLHandler(finalHandler))
	log.Println("Listening on port 3000 ......")
	http.ListenAndServe(":3000", nil)
}