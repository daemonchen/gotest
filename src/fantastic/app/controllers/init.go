package controllers

import (
	"flag"
	"github.com/huichen/sego"
	"github.com/jgraham909/revmgo"
	// "os"
	// "path/filepath"
	"runtime"
)

var Segmenter = sego.Segmenter{}
var dict = flag.String("dict", "src/fantastic/public/dictionary.txt", "词典文件")

func init() {
	revmgo.ControllerInit()
	// path, _ := os.Getwd()
	// 将线程数设置为CPU数
	runtime.GOMAXPROCS(runtime.NumCPU())

	// 初始化分词器
	Segmenter.LoadDictionary(*dict)
}
