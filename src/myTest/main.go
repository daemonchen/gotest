package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
)

/*

   结果：

   有时99
   有时100
*/

var counter_num uint32 = 0

const MAXSIZE = 100

func Count_num(ch chan uint32) {
	atomic.AddUint32(&counter_num, 1)
	ch <- counter_num
}

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	var chs [MAXSIZE]chan uint32

	for i := 0; i < MAXSIZE; i++ {
		chs[i] = make(chan uint32)
		go Count_num(chs[i])
	}

	for i, ch := range chs {
		print("-----", i, "\n")
		print(<-ch, "\n")
	}

	fmt.Println(counter_num)
}
