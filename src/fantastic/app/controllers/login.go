package controllers

import (
	"fantastic/app/models"
	// "fmt"
	"github.com/jgraham909/revmgo"
	"github.com/robfig/revel"
)

type Login struct {
	*revel.Controller
	revmgo.MongoController
}

type result struct {
	status string
	data   string
}

func (c Login) Index() revel.Result {
	return c.Render()

}

func (c Login) Login(username string, password string) revel.Result {
	responseJson := &result{}
	user := models.GetUserByName(c.MongoSession, username)
	if password == user.Password {
		c.Response.Status = 200
		c.Session["islogin"] = "true"
		return c.RenderJson(responseJson)
	} else {
		responseJson = &result{"caicaikana", "login failed"}
		c.Response.Status = 403
		c.Session["islogin"] = "false"
		return c.RenderJson(responseJson)

	}
}
