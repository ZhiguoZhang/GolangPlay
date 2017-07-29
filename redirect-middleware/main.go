package main

import (
	"github.com/ScruffyProdigy/Middleware/redirecter"
	"github.com/ScruffyProdigy/TheRack/httper"
)

func main() {
	conn := httper.HttpConnection(":3000")
	conn.Go(redirecter.Redirecter{Path: "https://docs.google.com/forms/d/e/1FAIpQLSe3sMjvlpw90JlIwEDFGmZACnPDXRjA8KGDbKGY3xjBsvYRig/viewform?usp=sf_link"})
}
