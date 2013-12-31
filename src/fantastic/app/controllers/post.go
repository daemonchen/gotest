package controllers

import (
	// "encoding/json"
	"crypto/md5"
	"fantastic/app/models"
	"github.com/jgraham909/revmgo"
	"github.com/robfig/revel"
	"io"
	// "labix.org/v2/mgo/bson"
	// "fmt"
	"math/rand"
	// "strconv"
	"time"
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
func (c *Post) generateSessionKey(data string) []byte {
	md5Key := md5.New()
	io.WriteString(md5Key, []byte(data))
	return md5Key.Sum("daemonchen")

}
func (c *Post) Index(stamp string) revel.Result {
	controllerName := "home"
	isLogin := c.Session["islogin"]

	randNum := rand.Int63n(time.Now().Unix())
	hashKey := c.generateSessionKey("this is my first hash session key")
	revel.WARN.Println("randNum:", randNum, hashKey)

	post := models.GetPostByStamp(c.MongoSession, stamp)
	comments := models.GetCommentsByStamp(c.MongoSession, stamp)

	return c.Render(controllerName, isLogin, post, comments)
}

func (c *Post) Update(stamp string, content string) revel.Result {
	revel.WARN.Println("commentData host:", c.Request.RemoteAddr)
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
	err := models.SaveComment(c.MongoSession, commentData)
	if err != nil {
		revel.WARN.Println("occur err when update:", err)
		return c.RenderJson(&BayesLearnResult{"insert comment failed", "failed comment"})
	}
	return c.RenderJson(responseJson)
}
