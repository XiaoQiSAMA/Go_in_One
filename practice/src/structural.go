package main

import "fmt"

func printArr(arr [5]int) {
	// 数组遍历
	for _, v := range arr {
		fmt.Println(v)
	}
}

func printArrPointer(arr *[5]int) {
	// 数组遍历
	arr[0] = 100
	for _, v := range arr {
		fmt.Println(v)
	}
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{2, 4, 6, 8, 10}

	var grid [4][5]int

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	printArr(arr1)
	// printArr(arr2)	cannot use arr2 (variable of type [3]int) as [5]int
	printArr(arr3)
	printArrPointer(&arr1)
	printArrPointer(&arr3)
}
