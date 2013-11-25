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
	Id            string `json:"id"`
	Title         string `json:"title"`
	Type          string `json:"type"`
	LatestNews    string `json:"latestNews"`
	UpdateTime    string `json:"updateTime"`
	HeadlineColor string `json:"headlineColor"`
	ChannelId     int    `json:"channelId"`
}
type Live struct {
	Name       string `json:"name"`
	Title      string `json:"title"`
	Id         int    `json:"id"`
	Type       string `json:"type"`
	Username   string `json:"username"`
	ServerId   int    `json:"serverId"`
	Nickname   string `json:"nickname"`
	UpdateTime int    `json:""updateTime`
	ZbStatus   int    `json:"zbStatus"`
	IsTop      int    `json:"isTop"`
	IsActive   bool   `json:"isActive"`
	LatestNews string `json:"latestNews"`
	Bulletin   string `json:"bulletin"`
}

type Badge struct {
	Master int `json:"master"`
	Note   int `json:"note"`
}

func (c Api) Update() revel.Result {
	// greeting := "Daemon"
	data := &Version{true, "orz", "6.3", true, "http://www.5800.com/ruanjian/app.apk"}
	return c.RenderJson(data)
}

func (c Api) Message() revel.Result {
	messageList := []interface{}{
		&Live{"AAA直播室", "AAA直播室", 5, "live", "jry001", 1, "森德金融研究所", 1318561963000, 1, 0, true, "春节关门行情，开心！", "<div>hello</div>"},
		&Message{"warning", "11f1", "warning", "hello", "2013-11-22T14:25:00.215Z", "#2293f5", 2},
		&Message{"chat", "我的投顾f", "chat", "hello", "2013-11-22T14:25:00.215Z", "#2293f5", 272},
		// &Message{"异动前哨", "qianshao", "上午大盘一下子就涨上去了又挂", "2013-11-22T14:25:00.215Z", "#2293f5", 273},
		&Message{"news_496", "系统公告f", "news", "关于白银最近不正常的通知", "2013-11-22T14:25:00.215Z", "#2293f5", 271},
		&Message{"strategyNews", "今日策略f", "strategyNews", "今天必须关注的经济指标", "123", "#2293f5", 275},
		&Message{"financialNews", "财经日历f", "financialNews", "欧元区8月份要倒闭", "123", "#f55200", 3},
		&Message{"hotNews_420", "热点聚焦f", "hotNews", "今天必须关注的经济指标", "123", "#f55200", 276},
		&Message{"news_498", "白银学堂f", "news", "白银学堂白银学堂白银学堂", "123", "#f55200", 321},
		&Message{"news_434", "快讯精灵f", "fastNews", "hello", "123", "#2293f5", 1},
		&Message{"news_495", "银江湖ff", "hotNews", "白银学堂白银学堂白银学堂", "123", "#f55200", 333},
		// &Message{"live", "天天涨停板f", "live", "天天涨停板", "2013-11-22T14:25:00.215Z", "#f55200", 333},
	}
	return c.RenderJson(messageList)
	// data := [] Message
}

func (c *Api) CheckBadgeInfo() revel.Result {
	result := &Badge{3, 4}
	return c.RenderJson(result)
}
