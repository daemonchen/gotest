package models

import (
	"encoding/json"
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
	return nil
}

func GetCommentsByStamp(s *mgo.Session, stamp string) comments {
	var comments comments
	getCommentCollection(s).Find(bson.M{"relativeStamp": stamp}).All(&comments)
	return comments
}
