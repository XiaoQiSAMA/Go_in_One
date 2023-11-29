# GoLang核心编程

## 1 语言基础

### 1-1 变量

#### 变量定义

```go
var a int
var s string

var a, b int = 3, 4

var a, b, c, s = 3, 4, true, "def"

// 只能在函数内使用
a, b, c, s := 3, 4, true, "def"
b = 5

// 并非全局变量，而是包内变量
var (
	aa = 3
	ss = "kkk"
	bb = true
)
```

#### 内建变量类型

* bool, string
* (u)int, (u)int8, (u)int16, (u)int64, uintptr
* btye(8bits), rune(32bits) 没有char只有rune
* float32, float64, complex64, complex128

#### 强制类型转换

```golang
var c int = int(math.Sqrt(float64(a*a + b*b)))
```

#### 常量

const 数值可以作为各种类型使用

```go
const filename = "abc.txt"
const a, b = 3, 4
```

#### 枚举

```go
// 普通枚举类型
const (
    cpp    = 1
    java   = 2
    python = 3
    golang = 4
)
// 自增枚举类型
const (
    // iota表示自增
    cpp = iota
    java
    python
    golang
)
```

### 1-2 结构语句

#### if ··· else ···

```go
if err != nil {
    ...
} else {
    ...
}

// contents, err为if内局部变量
if contents, err := ioutil.ReadFile(filename); err != nil {
    ...
} else {
    ...
}
```

#### switch

case内自动break,除非fallthrough

### for

```go
// 初始条件,退出条件,循环语句
for ; n > 0; n /= 2 {
    ...
}
// 仅有循环语句
for scanner.Scan() {
    ...
}
// 死循环
for {
    ...
}
```

### function

没有默认参数、可选参数

函数的同类型参数可以一同定义,函数的返回值为复数时须严格遵守数量(可以为返回值取名)

```go
func div(a, b int) (q, r int) {
    return a / b, a % b
}
```

函数参数传递

```go
// 匿名函数参数传递 apply(op func(int, int) int, a, b int) int
apply(
    func(i1, i2 int) int {
        return int(math.Pow(float64(i1), float64(i2)))
    }, 3, 4,
)
```

可变参数列表

```go
func sum(numbers ...int) int {
    ...
}
```

### 指针

Go语言只有值传递,通过指针来实现引用传递

```go
// *a, *b 表示指向的值
func swap(a, b *int) {
    *b, *a = *a, *b
}
```

### 数据结构

#### 数组

```go
var arr1 [5]int
arr2 := [3]int{1, 2, 3}
arr3 := [...]int{2, 4, 6, 8, 10}

var grid [4][5]int
```

range遍历

```go
for _, v := range arr3 {
    fmt.Println(v)
}
```

数组是值类型

```go
// cannot use arr2 (variable of type [3]int) as [5]int
func printArr(arr [5]int) {
    ...
}

// 数组的参数引用
func printArrPointer(arr *[5]int) {
    // 数组遍历
    arr[0] = 100
    ...
}
```

### 切片

slice是原本array的一个视图(指针)

![slice tips](images/slice.png)

slice底层实现依赖ptr,len,cap. 其中cap决定了slice能往后扩展,而不能往前扩展

![slice implement](images/slice_imply.png)

slice能够根据cap往后取,索引不能超过len取值

```go
fmt.Println("s2[4]=", s2[4])     // index out of range
fmt.Println("s2[4]=", s1[4:5])
```

#### append操作

```go
// 添加元素
s3 := append(s2, 10)
s4 := append(s3, 11)
s5 := append(s4, 12)
fmt.Println("s3, s4, s5 = ", s3, s4, s5)
// s4, s5添加的元素会覆盖原有arr元素值
// 超过arr的cap,系统则会分配新的更大cap的arr,把元素复制过去
// 后续会由垃圾回收机制处理
fmt.Println("arr = ", arr)

// output:
// slice of arr= [0 1 2 3 4 5 6 7]
// s3, s4, s5 =  [5 6 10] [5 6 10 11] [5 6 10 11 12]
// arr =  [0 1 2 3 4 5 6 10]
```

#### make创建

```go
s2 := make([]int, 16) // len = cap = 16
s3 := make([]int, 10, 32) // len = 10, cap = 32
```

#### copy

```go
copy(s2, s1) // s1拷给s2
```

#### delete

```go
// 把s2中0~2和4~end的slice加一起,来删除第3个元素
// s2[4:]...表示可变参数
s2 = append(s2[:3], s2[4:]...)

fmt.Println("Popping from front")
front := s2[0]
s2 = s2[1:]

fmt.Println("Popping from back")
tail := s2[len(s2)-1]
s2 = s2[:len(s2)-1]
```
