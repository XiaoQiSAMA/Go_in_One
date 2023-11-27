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

