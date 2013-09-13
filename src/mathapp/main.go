package main

import (
	"fmt"
	"mymath"
)

func main() {
	fmt.Printf("Hello,Sqrt(4)=%v\n", mymath.Sqrt(4))
	complex := 5 + 5i
	s := "hexixi"
	c := []byte(s)
	t := "a" + s[1:]
	rawS := `hexixi
  hahaha`
	fmt.Printf("complex data is %v", complex)
}
