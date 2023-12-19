package filelisting

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type userError string

func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string {
	return string(e)
}

const prefix = "/list/"

func HandlList(w http.ResponseWriter, r *http.Request) error {
	if strings.Index(r.URL.Path, prefix) != 0 {
		// 可以给用户看的err
		return userError("path must start with " + prefix)
	}
	path := r.URL.Path[len(prefix):]
	// fmt.Println(path)
	file, err := os.Open(path)
	if err != nil {
		// panic(err)
		// http.Error(w,
		// 	err.Error(),
		// 	http.StatusInternalServerError)
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	w.Write(all)
	return nil
}
