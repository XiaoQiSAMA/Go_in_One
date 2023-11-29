package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr[2:6]: ", arr[2:6])
	fmt.Println("arr[:6]: ", arr[:6])
	fmt.Println("arr[2:]: ", arr[2:])
	fmt.Println("arr[:]: ", arr[:])

	fmt.Println("Updare slice arr")
	updateSlice(arr[2:6])
	fmt.Println(arr)

	arr[2] = 2
	fmt.Println("Extending slice")
	s1 := arr[2:6]
	s2 := s1[3:5]
	// fmt.Println("s2[4]=", s2[4])		// index out of range
	fmt.Println("s2[4]=", s1[4:5])

	fmt.Println("s1=", s1)
	fmt.Println("s2=", s2)
	fmt.Printf("s1.len=%d, s1.cap=%d\n", len(s1), cap(s1))
	fmt.Printf("s2.len=%d, s2.cap=%d\n", len(s2), cap(s2))
	fmt.Printf("arr.len=%d, arr.cap=%d\n", len(arr), cap(arr))
	fmt.Println("slice of s1=", s1[:6])
	fmt.Println("slice of arr=", arr[:8])

	// 添加元素
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println("s3, s4, s5 = ", s3, s4, s5)
	// s4, s5添加的元素会覆盖原有arr元素值
	// 超过arr的cap,系统则会分配新的更大cap的arr,把元素复制过去
	// 后续会由垃圾回收机制处理
	fmt.Println("arr = ", arr)
}
