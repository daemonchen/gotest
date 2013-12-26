package controllers

import (
	"fantastic/app/models"
	"fmt"
	"github.com/jgraham909/revmgo"
	"github.com/robfig/revel"
	"labix.org/v2/mgo/bson"
	"strconv"
	"time"
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
	post := models.GetPostModel(bson.NewObjectId(), title, content, strconv.FormatInt(time.Now().Unix(), 10))
	post.Save(c.MongoSession)
	fmt.Println("post to save success")
	return c.RenderJson(responseJson)
}
//delete update when post update is finished
func (c *Edit) Update(stamp string, content string) revel.Result {
	responseJson := &BayesLearnResult{stamp, content}
	// post := models.GetPostModel(bson.NewObjectId(), title, content, strconv.FormatInt(time.Now().Unix(), 10))
	fmt.Println("post updated success")
	return c.RenderJson(responseJson)
}
