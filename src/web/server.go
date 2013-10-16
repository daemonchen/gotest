package main

import (
	"fmt"
	"labix.org/v2/mgo"
	// "labix.org/v2/mgo/bson"
	"html/template"
	"log"
	"net/http"
	"time"
	// "strings"
)

func dbTest(name string) {
	time := time.Now()
	session, err := mgo.Dial("localhost")
	if err != nil {
		fmt.Print("[db error]", err)
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("customer_behavior").C("people")
	err = c.Insert(&Person{name, "+55 53 8116 9639", time})
	if err != nil {
		fmt.Print("insert error", err)

		panic(err)
	}

	// result := Person{}
	// err = c.Find(bson.M{"name": "Ale"}).One(&result)
	// if err != nil {
	// 	fmt.Print(err)

	// 	panic(err)
	// }

	// fmt.Println("Phone:", result.Phone)
}
func logRequest(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	name := req.Form["name"]
	if len(name) > 0 {
		dbTest(name[0])
		t := template.New("mytemplate")
		r, _ := t.Parse("hello,{{.}}!")
		r.Execute(res, name[0])
	}
}

type Person struct {
	Name  string
	Phone string
	Stamp time.Time
}

func main() {
	http.HandleFunc("/", logRequest)
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal("error on server", err)
	}
	fmt.Print("start server on port 9999")

}
