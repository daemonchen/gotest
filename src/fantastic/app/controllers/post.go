package controllers

import (
	"fantastic/app/models"
	"github.com/jgraham909/revmgo"
	"github.com/robfig/revel"
	// "fmt"
	// "strconv"
	// "time"
)

type Post struct {
	*revel.Controller
	revmgo.MongoController
}

func (c *Post) Index(stamp string) revel.Result {
	controllerName := "home"
	isLogin := c.Session["islogin"]

	post := models.GetPostByStamp(c.MongoSession, stamp)
	// time4int64, _ := strconv.ParseInt(post.Stamp, 10, 64)
	// timeUtc := time.Unix(time4int64, 0)
	// const layout = "Jan 2, 2006 at 3:04pm (MST)"
	// post.Stamp = timeUtc.Format(layout)
	revel.WARN.Println("query post success")
	return c.Render(controllerName, isLogin, post)
// }

func (c *Post) Update(stamp string, content string) revel.Result {
	responseJson := &BayesLearnResult{stamp, "success update"}
	err := models.UpdatePost(c.MongoSession, stamp, content)
	if err != nil {
		revel.WARN.Println("occur err when update:", err)
	}
	// revel.WARN.Println("post updated success")
	return c.RenderJson(responseJson)
// }
