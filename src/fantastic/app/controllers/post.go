package controllers

import (
	// "encoding/json"
	"crypto/md5"
	"fantastic/app/models"
	"github.com/jgraham909/revmgo"
	"github.com/robfig/revel"
	// "github.com/robfig/revel/cache"
	"io"
	// "labix.org/v2/mgo/bson"
	// "fmt"
	"math/rand"
	"strconv"
	"time"
)

type Post struct {
	*revel.Controller
	revmgo.MongoController
}

var CommentCache = make(map[string]string)

func (c *Post) generateSessionKey() []byte {
	md5Key := md5.New()
	io.WriteString(md5Key, "this is my first hash session key")
	return md5Key.Sum([]byte("daemonchen"))

}
func (c *Post) Index(stamp string) revel.Result {
	controllerName := "home"
	isLogin := c.Session["islogin"]

	randNum := rand.Int63n(time.Now().Unix())
	hashKey := c.generateSessionKey()
	c.Session[string(hashKey[:])] = strconv.FormatInt(randNum, 10)
	CommentCache[strconv.FormatInt(randNum, 10)] = "true"

	post := models.GetPostByStamp(c.MongoSession, stamp)
	if len(post.Content) > 0 {
		comments := models.GetCommentsByStamp(c.MongoSession, stamp)
		return c.Render(controllerName, isLogin, post, comments)
	} else {
		return c.Redirect(App.Index)
	}
}

func (c *Post) Delete(stamp string) revel.Result {
	// controllerName := "home"
	if isLogin := c.Session["islogin"]; isLogin == "true" {
		err := models.DeletePost(c.MongoSession, stamp)
		if err != nil {
			panic(err)
		}
	}
	return c.RenderJson(&BayesLearnResult{stamp, "success update"})
}

func (c *Post) Update(stamp string, content string) revel.Result {
	revel.WARN.Println("commentData host:", c.Request.RemoteAddr)
	responseJson := &BayesLearnResult{"delete", "success delete"}
	err := models.UpdatePost(c.MongoSession, stamp, content)
	if err != nil {
		revel.WARN.Println("occur err when update:", err)
	}
	revel.WARN.Println("post updated success")
	return c.RenderJson(responseJson)
}

func (c *Post) clearCommentCacheValue() {
	hashKey := c.generateSessionKey()
	CommentCache[c.Session[string(hashKey[:])]] = "diff"
}
func (c *Post) AddComment(commentData string) revel.Result {
	hashKey := c.generateSessionKey()
	if CommentCache[c.Session[string(hashKey[:])]] == "diff" {
		revel.WARN.Println("CommentCache is diff")
		c.Response.Status = 403
		return c.RenderJson(&BayesLearnResult{"failed", "you can't comment now, please refresh page and wait for a moment"})
	}
	err := models.SaveComment(c.MongoSession, commentData)
	if err != nil {
		revel.WARN.Println("occur err when update:", err)
		c.Response.Status = 403
		return c.RenderJson(&BayesLearnResult{"failed", "insert comment failed"})
	}
	c.clearCommentCacheValue()
	c.Response.Status = 200
	return c.RenderJson(&BayesLearnResult{"success", "insert comment success"})

}
