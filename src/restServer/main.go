package main

import (
	"fmt"
	"github.com/drone/routes"
	"net/http"
)

func getUser(writer http.ResponseWriter, reader *http.Request) {
	params := reader.URL.Query()
	uid := params.Get(":uid")
	fmt.Fprintf(writer, "you are get user: %s", uid)
}

func updateUser(writer http.ResponseWriter, reader *http.Request) {
	params := reader.URL.Query()
	uid := params.Get(":uid")
	fmt.Fprintf(writer, "you are update user: %s", uid)
}

func deleteUser(writer http.ResponseWriter, reader *http.Request) {
	params := reader.URL.Query()
	uid := params.Get(":uid")
	fmt.Fprintf(writer, "you are delete user: %s", uid)
}

func addUser(writer http.ResponseWriter, reader *http.Request) {
	params := reader.URL.Query()
	uid := params.Get(":uid")
	fmt.Fprintf(writer, "you are add user: %s", uid)
}

func main() {
	mux := routes.New()
	mux.Get("/user/:uid", getUser)
	mux.Post("/user/:uid", updateUser)
	mux.Del("/user/:uid", deleteUser)
	mux.Put("/user/", addUser)

	http.Handle("/", mux)
	err := http.ListenAndServe(":1988", nil)
	if err != nil {
		fmt.Printf("err occur when server start: %s", err)
	}
	fmt.Print("server start at:1988")
}
