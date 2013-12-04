package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	// "mymath"
	. "github.com/jbrukh/bayesian"
)

const (
	Good Class = "Good"
	Bad  Class = "Bad"
)

// var classifier *Classifier
func readTemplate(file string) (stringArray []string) {
	byteArray, _ := ioutil.ReadFile(file)
	stringResult := string(byteArray[:])
	fmt.Println("stringResult", stringResult)
	stringArray = strings.Split(stringResult, ",")
	return
}

func bayesLearn() {
	classifier := NewClassifier(Good, Bad)
	goodStuff := readTemplate("good.txt")
	badStuff := readTemplate("bad.txt")
	classifier.Learn(goodStuff, Good)
	classifier.Learn(badStuff, Bad)
	writer := bytes.NewBuffer(nil)
	classifier.WriteTo(writer)
	ioutil.WriteFile("class.txt", writer.Bytes(), os.ModeAppend|os.ModePerm)
}

// func complexTest() {
// 	fmt.Printf("Hello,Sqrt(4)=%v\n", mymath.Sqrt(4))
// 	complex := 5 + 5i
// 	s := "daemon"
// 	c := []byte(s)
// 	t := "a" + s[1:]
// 	fmt.Printf("complex data is %v", complex)
// }

func logScores() {
	classifier, _ := NewClassifierFromFile("class.txt")
	scores, likely, _ := classifier.LogScores([]string{"tall", "girl"})
	fmt.Println("--->>>:", scores, likely)
}
func main() {
	bayesLearn()
	logScores()
}
