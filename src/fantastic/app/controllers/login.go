package controllers

import (
	"fantastic/app/models"
	"fmt"
	"github.com/jgraham909/revmgo"
	"github.com/robfig/revel"
)

type Login struct {
	*revel.Controller
	revmgo.MongoController
}

func (c Login) Index() revel.Result {
	return c.Render()

}

func (c Login) Login(userName string, password string) revel.Result {
	fmt.Println(userName, password)
	user := models.GetUserByName(c.MongoSession, "daemon")
	return c.RenderJson(user)
}
