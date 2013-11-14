package controllers

import (
	// "encoding/json"
	// "fantastic/app/models"
	// "fmt"
	"github.com/jgraham909/revmgo"
	"github.com/robfig/revel"
)

type Api struct {
	*revel.Controller
	revmgo.MongoController
}

type Version struct {
	ForceUpdate bool   `json:"forceUpdate"`
	Description string `json:"description"`
	VersionId   string `json:"versionId"`
	CanUpdate   bool   `json:"canUpdate"`
	AppUrl      string `json:"appUrl"`
}

func (c Api) Update() revel.Result {
	// greeting := "Daemon"
	data := &Version{true, "orz", "6.3", true, "http://www.5800.com/ruanjian/app.apk"}
	return c.RenderJson(data)
}
