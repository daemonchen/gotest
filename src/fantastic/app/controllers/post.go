package controllers

import (
	// "encoding/json"
	"fantastic/app/models"
	"github.com/jgraham909/revmgo"
	"github.com/robfig/revel"
	// "labix.org/v2/mgo/bson"
	// "fmt"
	// "strconv"
	// "time"
)

type Post struct {
	*revel.Controller
	revmgo.MongoController
}

// type Comment struct {
// 	// Id            bson.ObjectId `json:"_id,omitempty"`
// 	RelativeStamp string `json:"relativeStamp"`
// 	UserName      string `json:"userName"`
// 	UserEmail     string `json:"userEmail"`
// 	CommentText   string `json:"commentText"`
// 	CommentTime   string `json:"commentTime"`
// }

func (c *Post) Index(stamp string) revel.Result {
	controllerName := "home"
	isLogin := c.Session["islogin"]

	post := models.GetPostByStamp(c.MongoSession, stamp)
	comments := models.GetCommentsByStamp(c.MongoSession, stamp)
	revel.WARN.Println("query post success")
	return c.Render(controllerName, isLogin, post, comments)
}

func (c *Post) Update(stamp string, content string) revel.Result {
	responseJson := &BayesLearnResult{stamp, "success update"}
	err := models.UpdatePost(c.MongoSession, stamp, content)
	if err != nil {
		revel.WARN.Println("occur err when update:", err)
	}
	revel.WARN.Println("post updated success")
	return c.RenderJson(responseJson)
}

func (c *Post) AddComment(commentData string) revel.Result {
	responseJson := &BayesLearnResult{"success comment", "success comment"}
	// var comment Comment
	// err := json.Unmarshal([]byte(commentData), &comment)
	revel.WARN.Println("commentData host:", c.Request)
	err := models.SaveComment(c.MongoSession, commentData)
	if err != nil {
		revel.WARN.Println("occur err when update:", err)
		return c.RenderJson(&BayesLearnResult{"insert comment failed", "failed comment"})
	}
	return c.RenderJson(responseJson)
}
