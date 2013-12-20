package controllers

import (
	// "fantastic/app/models"
	// "fmt"
	"github.com/jgraham909/revmgo"
	"github.com/robfig/revel"
)

type Edit struct {
	*revel.Controller
	revmgo.MongoController
}

func (c Edit) Index() revel.Result {
	if c.Session["islogin"] != "true" {
		return c.Redirect(Login.Index)
	}
	controllerName := "edit"
	return c.Render(controllerName)

}
func (c *Edit) Post(title string, content string) revel.Result {
	responseJson := &BayesLearnResult{"success", "article saved success"}
	return c.RenderJson(responseJson)
}
