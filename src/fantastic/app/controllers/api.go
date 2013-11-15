package controllers

import (
	// "encoding/json"
	// "fantastic/app/models"
	// "fmt"
	"github.com/jgraham909/revmgo"
	"github.com/robfig/revel"
)

type Api struct {
	*revel.Controller
	revmgo.MongoController
}

type Version struct {
	ForceUpdate bool   `json:"forceUpdate"`
	Description string `json:"description"`
	VersionId   string `json:"versionId"`
	CanUpdate   bool   `json:"canUpdate"`
	AppUrl      string `json:"appUrl"`
}

type Message struct {
	Title         string `json:"title"`
	Type          string `json:"type"`
	LatestNews    string `json:"latestNews"`
	UpdateTime    int    `json:"updateTime"`
	HeadlineColor string `json:"headlineColor"`
	ChannelId     int    `json:"channelId"`
}

func (c Api) Update() revel.Result {
	// greeting := "Daemon"
	data := &Version{true, "orz", "6.3", true, "http://www.5800.com/ruanjian/app.apk"}
	return c.RenderJson(data)
}

func (c Api) Message() revel.Result {
	messageList := []*Message{
		&Message{"预警通知", "warning", "hello", 123, "#2293f5", 2},
		&Message{"我的投顾", "chat", "hello", 123, "#2293f5", 272},
		&Message{"异动前哨", "qianshao", "上午大盘一下子就涨上去了又挂", 123, "#2293f5", 273},
		&Message{"系统公告", "news", "关于白银最近不正常的通知", 123, "#2293f5", 271},
		&Message{"今日策略", "strategyNews", "今天必须关注的经济指标", 123, "#2293f5", 275},
		&Message{"财经日历", "financialNews", "欧元区8月份要倒闭", 123, "#f55200", 3},
		&Message{"热点聚焦", "hotNews", "今天必须关注的经济指标", 123, "#f55200", 276},
		&Message{"白银学堂", "news", "白银学堂白银学堂白银学堂", 123, "#f55200", 321},
		&Message{"快讯精灵", "warning", "hello", 123, "#2293f5", 1},
		&Message{"银江湖", "hotNews", "白银学堂白银学堂白银学堂", 123, "#f55200", 333},
	}
	return c.RenderJson(messageList)
	// data := [] Message
}
