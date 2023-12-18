package main

import (
	"bufio"
	"fmt"
	fib "learngo/functional"
	"os"
)

func tryDefer() {
	// defer存入栈中
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	panic("error occurred!")
	fmt.Println(4)
}

func writerFile(filename string) {
	// file, err := os.Create(filename)
	file, err := os.OpenFile(
		filename, os.O_EXCL|os.O_CREATE, 0666,
	)
	// // 自定义err
	// err = errors.New("This is a custom error!")
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s, %s, %s\n",
				pathError.Op,
				pathError.Path,
				pathError.Err)
		}
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}

}

func main() {
	// tryDefer()
	writerFile("fib.txt")
}
