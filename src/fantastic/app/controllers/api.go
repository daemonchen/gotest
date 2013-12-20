package controllers

import (
	// "encoding/json"
	// "fantastic/app/models"
	"bytes"
	"fmt"
	. "github.com/jbrukh/bayesian"
	"github.com/jgraham909/revmgo"
	"github.com/robfig/revel"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	// "strings"
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

type SegmentField struct {
	Text string `json:"text"`
	Pos  string `json:"pos"`
}

type JsonResponse struct {
	Segments []*SegmentField `json:"segments"`
}

type BayesLearnResult struct {
	Status string `json:"status"`
	Data   string `json:"data"`
}

type BayesScore struct {
	Good   string `json:"good"`
	Bad    string `json:"bad"`
	Likely string `json:"likely"`
	Strict bool   `json:"strict"`
}

const (
	Good = "Good"
	Bad  = "Bad"
)

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

func bayesLearn(text []string, class Class) *BayesLearnResult {
	var classifier *Classifier
	wd, _ := os.Getwd()
	fmt.Println(">>>>>>os wd", wd)
	classifier, err := NewClassifierFromFile("class.txt")
	if err != nil {
		classifier = NewClassifier(Good, Bad)
	}
	classifier.Learn(text, class)
	writer := bytes.NewBuffer(nil)
	classifier.WriteTo(writer)
	ioutil.WriteFile("class.txt", writer.Bytes(), os.ModeAppend|os.ModePerm)
	responseJson := &BayesLearnResult{"success", "bayes learn success"}
	fmt.Println("classifier:", classifier)
	return responseJson
}

func (c *Api) LogScore(text string) revel.Result {
	classifier, err := NewClassifierFromFile("class.txt")
	if err != nil {
		responseJson := &BayesLearnResult{"error", "get score failed"}
		return c.RenderJson(responseJson)
	} else {
		// fmt.Println(">>>>>>>>>classifier", classifier)
		scores, likely, strict := classifier.LogScores(strings.Fields(text))
		fmt.Println("LogScore:", scores, likely)
		responseJson := &BayesScore{strconv.FormatFloat(scores[0], 'f', 2, 64), strconv.FormatFloat(scores[1], 'f', 2, 64), string(classifier.Classes[likely]), strict}
		fmt.Println("responseJson:", responseJson)
		return c.RenderJson(responseJson)
	}
}

func (c *Api) Segment(text string) revel.Result {
	// 分词
	segments := Segmenter.Segment([]byte(text))

	// 整理为输出格式
	ss := []*SegmentField{}
	for _, segment := range segments {
		ss = append(ss, &SegmentField{Text: segment.Token().Text(), Pos: segment.Token().Pos()})
	}
	response := &JsonResponse{Segments: ss}
	// fmt.Println(">>>>>>>>>>>>", response)
	return c.RenderJson(response)
}

func (c *Api) Bayes(category string, text string) revel.Result {
	// space := []byte{' '}
	parts := strings.Fields(text)
	result := bayesLearn(parts, Class(category))
	fmt.Println("result", result)

	return c.RenderJson(result)
}
