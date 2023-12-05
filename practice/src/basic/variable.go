package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func euler() {
	// fmt.Println(
	// 	cmplx.Pow(math.E, 1i*math.Pi) + 1,
	// )

	// fmt.Println(cmplx.Exp(1i*math.Pi) + 1)

	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
}

func triangle() {
	var a, b int = 3, 4
	var c int = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func consts() {
	const filename = "abc.txt"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enums() {
	// const (
	// 	cpp    = 1
	// 	java   = 2
	// 	python = 3
	// 	golang = 4
	// )
	const (
		// iota表示自增
		cpp = iota
		java
		python
		golang
	)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, java, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}
