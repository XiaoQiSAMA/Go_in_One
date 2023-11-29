package main

import "fmt"

func printSlice(slice []int) {
	fmt.Printf("%v, len=%d, cap=%d\n",
		slice, len(slice), cap(slice))
}

func main() {
	var s []int // Zero value for slice is nil

	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, i*2+1)
	}

	fmt.Println(s)

	s1 := []int{2, 4, 6, 8}
	printSlice(s1)

	s2 := make([]int, 16) // len = cap = 16

	s3 := make([]int, 10, 32) // len = 10, cap = 32

	printSlice(s2)
	printSlice(s3)

	fmt.Println("Copying slice")
	copy(s2, s1) // s1拷给s2
	printSlice(s2)

	fmt.Println("Deleting slice")
	// 把s2中0~2和4~end的slice加一起,来删除第3个元素
	// s2[4:]...表示可变参数
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	fmt.Println("Popping from front")
	front := s2[0]
	s2 = s2[1:]

	fmt.Println("Popping from back")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]

	fmt.Println(front, tail)
	printSlice(s2)
}
