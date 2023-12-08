package main

import (
	"fmt"
	testCode "learngo/testing"
)

func getRetriever() retriever {
	return testCode.Retriever{}
}

type retriever interface {
	Get(string) string
}

func main() {
	var r retriever = getRetriever()
	fmt.Println(r.Get("https://www.imooc.com"))
}
