package controllers

import (
	// "encoding/json"
	"fantastic/app/models"
	"fmt"
	"github.com/jgraham909/revmgo"
	"github.com/robfig/revel"
	// "labix.org/v2/mgo/bson"
	"strconv"
	"time"
)

type App struct {
	*revel.Controller
	revmgo.MongoController
}
type posts []interface{}

func (c App) Index() revel.Result {
	controllerName := "home"
	posts := models.GetAllPosts(c.MongoSession)
	for _, post := range posts {
		time4int64, _ := strconv.ParseInt(post.Stamp, 10, 64)
		timeUtc := time.Unix(time4int64, 0)
		const layout = "Jan 2, 2006 at 3:04pm (MST)"
		fmt.Println("post:", post)
		fmt.Println("post time:", timeUtc.Format(layout))

	}
	return c.Render(controllerName, posts)
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
