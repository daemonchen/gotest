package main

import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
	"net/http"
	// "strings"
)

func logRequest(res http.ResponseWriter, req *http.Request) {
	fmt.Print("aa")
	req.ParseForm()
	fmt.Println(req.Form)
	fmt.Println(req.URL.Path)
	fmt.Fprintf(res, "test")
}

type Person struct {
	Name  string
	Phone string
}

func main() {
	http.HandleFunc("/", logRequest)
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal("error on server", err)
	}
	fmt.Print("start server on port 9999")

	session, err := mgo.Dial("http://192.168.27.108:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("customer_behavior").C("people")
	err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
		&Person{"Cla", "+55 53 8402 8510"})
	if err != nil {
		panic(err)
	}

	result := Person{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		panic(err)
	}

	fmt.Println("Phone:", result.Phone)
}
