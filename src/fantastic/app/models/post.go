package models

import (
	// "encoding/json"
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type Post struct {
	Id      bson.ObjectId `bson:"_id,omitempty"`
	Title   string        `bson:"title"`
	Content string        `bson:"content"`
	Stamp   string        `bson:"stamp"`
}

type posts []*Post

func getPostsCollection(s *mgo.Session) *mgo.Collection {
	return s.DB("fantastic").C("posts")
}
func (b *Post) Save(s *mgo.Session) error {
	err := getPostsCollection(s).Insert(b)
	if err != nil {
		fmt.Println("save post occur error:", err)
	}
	return err
}

func GetPostModel(id bson.ObjectId, title string, content string, stamp string) *Post {
	post := &Post{id, title, content, stamp}
	return post
}
func GetAllPosts(s *mgo.Session) posts {
	var posts posts
	getPostsCollection(s).Find(nil).All(&posts)
	return posts
}
