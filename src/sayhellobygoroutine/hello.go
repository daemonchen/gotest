package main

import (
	"fmt"
	// "os/exec"
	"runtime"
)

func hello(s string) {
	for i := 0; i < 3; i++ {
		runtime.Gosched()
		fmt.Println(s)
		fmt.Println(runtime.NumCPU())
	}
}
func iteratorTest() {
	m := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
	//slice
	s := make([]int, len(m))
	i := 0
	for k, _ := range m {
		s[i] = k
		i++
	}
	fmt.Println("iterator test", s, len(m))
	// fmt.Println(s)
}
func main() {
	go hello("1")
	hello("2")
	// cmd := exec.Command("shutdown", "-f")
	// cmd.Run()
	hello("3")
	iteratorTest()
}
