package models

import (
	// "encoding/json"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type Manager struct {
	Id       bson.ObjectId `bson:"_id,omitempty"`
	UserName string        `bson:"username"`
	PassWord string        `bson:"password"`
}

type managers []*Manager

func getManageeCollection(s *mgo.Session) *mgo.Collection {
	return s.DB("fantastic").C("adminMailAccount")
}

func GetManager(s *mgo.Session) *Manager {
	var manager *Manager
	getManageeCollection(s).Find(nil).One(&manager)
	return manager
}
