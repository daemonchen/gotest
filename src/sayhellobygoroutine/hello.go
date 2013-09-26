package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func hello(s string) {
	for i := 0; i < 3; i++ {
		runtime.Gosched()
		fmt.Println(s)
		fmt.Println(runtime.NumCPU())
	}
}

func main() {
	go hello("1")
	hello("2")
	cmd := exec.Command("shutdown", "-f")
	cmd.Run()
	// hello("3")
}
