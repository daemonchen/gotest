package main

import (
	"encoding/json"
	"fmt"
	"github.com/drone/routes"
	"net/http"
	"restServer/logs"
)

type Badge struct {
	Live   int `json:"live"`
	Master int `json:"master"`
	Chat   int `json:"chat"`
	Note   int `json:"note"`
}
type Result struct {
	Success bool   `json:"success"`
	Data    *Badge `json:"data"`
}

func getUser(writer http.ResponseWriter, reader *http.Request) {
	params := reader.URL.Query()
	uid := params.Get(":uid")
	logs.Logger.Info("test from go")
	go logs.Logger.Critical("test from go")
	fmt.Fprintf(writer, "you are get user: %s", uid)
}

func getBadgeInfo(writer http.ResponseWriter, reader *http.Request) {
	reader.ParseForm()
	badges := &Badge{1, 2, 3, 4}
	data := Result{Success: true, Data: badges}
	result, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	writer.Write(result)
	fmt.Println("result: ", string(result))
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
	mux.Get("/liveRoom/checkBadgeInfo.do", getBadgeInfo)
	mux.Post("/user/:uid", updateUser)
	mux.Del("/user/:uid", deleteUser)
	mux.Put("/user/", addUser)

	http.Handle("/", mux)
	fmt.Println("server start at:1988")

	if err := http.ListenAndServe(":1988", nil); err != nil {
		fmt.Printf("err occur when server start: %s", err)
	}
}
