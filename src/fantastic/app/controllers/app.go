package controllers

import (
	// "encoding/json"
	"fantastic/app/models"
	"fmt"
	"github.com/jgraham909/revmgo"
	"github.com/robfig/revel"
)

type App struct {
	*revel.Controller
	revmgo.MongoController
}

type Badge struct {
	Master int `json:"master"`
	Note   int `json:"note"`
}

func (c App) Index() revel.Result {
	greeting := "Daemon"
	return c.Render(greeting)
}

func (c App) Hello(myName string) revel.Result {
	c.Validation.Required(myName).Message("please input your name")
	c.Validation.MinSize(myName, 3).Message("name is not long enough")
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Index)
	}
	return c.Render(myName)
}

func (c App) Users() revel.Result {
	result := models.GetUserByName(c.MongoSession, "daemon")
	return c.Render(result)
}

func (c *App) AllUsers() revel.Result {
	data := models.GetAllUsers(c.MongoSession)
	fmt.Println(">>>>>>all users in action:", data)
	return c.RenderJson(data)
}
