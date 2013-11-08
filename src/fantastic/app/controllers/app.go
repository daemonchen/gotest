package controllers

import (
	"fantastic/app/models"
	"github.com/jgraham909/revmgo"
	"github.com/robfig/revel"
)

type App struct {
	*revel.Controller
	revmgo.MongoController
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
