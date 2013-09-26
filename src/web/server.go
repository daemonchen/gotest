package main

import (
	"fmt"
	"log"
	"net/http"
	// "strings"
)

func logRequeset(req http.ResponseWriter, res *http.Request) {
	res.ParseForm()
	fmt.Println(res.Form)
	fmt.Println(res.URL.Path)
	fmt.Fprintf(req, "test")
}

func main() {
	http.HandleFunc("/", logRequeset)
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal("error on server", err)
	}
}
