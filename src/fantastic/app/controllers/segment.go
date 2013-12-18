package controllers

import (
	// "fmt"
	"github.com/jgraham909/revmgo"
	"github.com/robfig/revel"
)

type Segment struct {
	*revel.Controller
	revmgo.MongoController
}

func (c *Segment) Index() revel.Result {
	controllerName := "segment"
	return c.Render(controllerName)
}
