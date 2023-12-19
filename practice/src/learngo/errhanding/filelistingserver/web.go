package main

import (
	"fmt"
	"learngo/errhanding/filelistingserver/filelisting"
	"log"
	"net/http"
	"os"
)

type appHandler func(writer http.ResponseWriter,
	request *http.Request) error

// 针对appHandler函数做出错处理
func errWrapper(handler appHandler) func(
	http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic: %v", r)
				http.Error(w,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()

		err := handler(w, r)
		if err != nil {
			log.Printf("Error handling request: %s",
				err.Error())
			fmt.Printf("err type: %T\n", err)
			if userErr, ok := err.(userError); ok {
				http.Error(w,
					userErr.Message(),
					http.StatusBadRequest)
				return
			}
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(w,
				http.StatusText(code), code)
		}
	}
}

type userError interface {
	error
	Message() string
}

func main() {
	http.HandleFunc("/", errWrapper(filelisting.HandlList))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
