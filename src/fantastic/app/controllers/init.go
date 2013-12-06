package controllers

import (
	"flag"
	"fmt"
	"github.com/huichen/sego"
	"github.com/jgraham909/revmgo"
	"os"
	// "path/filepath"
	"runtime"
)

var Segmenter = sego.Segmenter{}
var dict = flag.String("dict", "src/fantastic/public/dictionary.txt", "词典文件")

func init() {
	revmgo.ControllerInit()
	// 将线程数设置为CPU数
	path, _ := os.Getwd()
	fmt.Println("server start at >>>>>>", path)
	runtime.GOMAXPROCS(runtime.NumCPU())

	// 初始化分词器
	go Segmenter.LoadDictionary(*dict)
}
