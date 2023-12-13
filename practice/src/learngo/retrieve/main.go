package main

import (
	"fmt"
	"learngo/retrieve/mock"
	"learngo/retrieve/real"
	"time"
)

const url = "http://www.imooc.com"

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string,
		form map[string]string) string
}

func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(url,
		map[string]string{
			"name":   "ccmouse",
			"course": "golang",
		})
}

// 组合Retriever与Poster
type RetrieverPoster interface {
	Retriever
	Poster
}

// 组合方法的使用
func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked imooc.com",
	})
	return s.Get(url)
}

func main() {
	var r Retriever

	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	fmt.Printf("%T %v\n", r, r)

	// Type assertion: 取出r中的值
	realRetriever := r.(*real.Retriever)
	fmt.Println(realRetriever.UserAgent)

	r = &mock.Retriever{Contents: "This is a fake imooc.com"}
	// fmt.Printf("%T %v\n", r, r)
	fmt.Println("Inspecting", r)

	if mockRetriever, ok := r.(mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}

	fmt.Println("Try a seesion")
	s := &mock.Retriever{Contents: "This is a fake imooc.com"}
	fmt.Println(session(s))

	// fmt.Println(download(r))
}
