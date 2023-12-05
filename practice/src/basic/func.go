package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		// return a / b
		q, _ := div(a, b)
		return q, nil
	default:
		// panic("unsupported op: " + op)
		return 0, fmt.Errorf("unsupported op: " + op)
	}
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args "+"(%d, %d)\n", opName, a, b)

	return op(a, b)
}

func div(a, b int) (q, r int) {
	return a / b, a % b
}

// 可变参数列表
func sum(numbers ...int) int {
	sum := 0
	for i := range numbers {
		sum += numbers[i]
	}
	return sum
}

// 指针
func swap(a, b *int) {
	*b, *a = *a, *b
}

func main() {
	if result, err := eval(3, 4, "x"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
	q, r := div(13, 3)
	fmt.Println(q, r)

	// 函数参数传递
	fmt.Println(apply(
		func(i1, i2 int) int {
			return int(math.Pow(float64(i1), float64(i2)))
		}, 3, 4,
	))

	fmt.Println(sum(1, 2, 3, 4, 5))

	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)
}
