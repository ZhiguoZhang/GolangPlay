package main

import (
    "log"
    "net/http"
    "time"
)

//Code example for http://www.alexedwards.net/blog/a-recap-of-request-handling

func timeHandler(format string) http.Handler {
    fn := func(w http.ResponseWriter, r *http.Request) {
        tm := time.Now().Format(format)
        w.Write([]byte("The time is: " + tm))
    }
    return http.HandlerFunc(fn)
}

func main() {


    th1123 := timeHandler(time.RFC1123)

    http.Handle("/time/rfc1123", th1123)

    th3339 := timeHandler(time.RFC3339)

    http.Handle("/time/rfc3339", th3339)

    log.Println("Listening.....")
    http.ListenAndServe(":3000", nil)
}