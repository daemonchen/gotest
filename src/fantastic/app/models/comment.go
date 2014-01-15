package models

import (
	"encoding/json"
	. "fantastic/app/lib/email"
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type Comment struct {
	Id            bson.ObjectId `bson:"_id,omitempty"`
	RelativeStamp string        `bson:"relativeStamp"`
	UserName      string        `bson:"userName"`
	UserEmail     string        `bson:"userEmail"`
	CommentText   string        `bson:"commentText"`
	CommentTime   string        `bson:"commentTime"`
}

type comments []*Comment

func getCommentCollection(s *mgo.Session) *mgo.Collection {
	return s.DB("fantastic").C("comments")
}

func sendMailToManager(m *Manager, c *Comment) {
	emailTitle := fmt.Sprintf("Here is the new comment from %s\n", c.UserName)
	emailContent := fmt.Sprintf(" The comment is:\n %s You can scan the post in http://115.29.47.52/post/index?stamp=%s ", c.CommentText, c.RelativeStamp)
	Mail(m.UserName, m.PassWord, emailTitle, emailContent)
}
func SaveComment(s *mgo.Session, commentRaw string) error {
	comment := &Comment{}
	jsonErr := json.Unmarshal([]byte(commentRaw), comment)
	if jsonErr != nil {
		panic(jsonErr)
		return jsonErr
	}
	err := getCommentCollection(s).Insert(comment)
	if err != nil {
		panic(err)
		return err
	}
	manager := GetManager(s)
	go sendMailToManager(manager, comment)
	return nil
}

func GetCommentsByStamp(s *mgo.Session, stamp string) comments {
	var comments comments
	getCommentCollection(s).Find(bson.M{"relativeStamp": stamp}).All(&comments)
	return comments
}
