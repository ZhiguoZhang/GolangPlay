package main

import (
	"fmt"
	//"github.com/ScruffyProdigy/Middleware/redirecter"
	//"github.com/ScruffyProdigy/TheRack/httper"
	"io/ioutil"
	//"log"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://akademikerchor.berlin")
	check(err)
	body, err := ioutil.ReadAll(resp.Body)
	check(err)
	fmt.Println(len(body))

}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
