package main

import "fmt"

// var aa = 3
// var ss = "kkk"
// var bb = true

// 并非全局变量，而是包内变量
var (
	aa = 3
	ss = "kkk"
	bb = true
)

// err: var bb := 1

func variableZeroValue() {
	var a int
	var s string
	// %s是字符串，%q是带引号的字符串
	fmt.Printf("%d, %q\n", a, s)
}

func variableInitValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	// 只能在函数内使用
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}

func main() {
	fmt.Println("hello world!")
	variableZeroValue()
	variableInitValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, ss, bb)

	euler()
	triangle()
	consts()
	enums()
}
