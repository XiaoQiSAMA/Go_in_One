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

### 1-3 指针

Go语言只有值传递,通过指针来实现引用传递

```go
// *a, *b 表示指向的值
func swap(a, b *int) {
    *b, *a = *a, *b
}
```

### 1-4 数据结构

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

### 1-5 切片

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

### map

· map[k]v  map[k1]map[k2]v

map的key:

1. map使用哈希表,必须可以比较相等
2. 除了slice,map,function的内建类型都可以为key
3. Struct类型不包含上述字段,也可作为key

```go
// 创建
// hash map
m := map[string]string{
    "name":    "ccmouse",
    "course":  "golang",
    "site":    "imooc",
    "quality": "notbad",
}

m2 := make(map[string]int) // m2 == empty map

var m3 map[string]int // m3 == nil

fmt.Println("Traversing map")
for k, v := range m {
    fmt.Println(k, v)
}

fmt.Println("Getting values")
courseName := m["course"]

if causeName, ok := m["cause"]; ok {
    fmt.Println(causeName)
} 

fmt.Println("Deleting values")

delete(m, "name")
```

### rune类型(Strings操作)

rune相当于go的char

1. 使用range遍历pos,rune对
2. 使用utf8.RuneCountInString获得字符数量
3. 使用len获得字节长度
4. 使用[]byte获得字节

#### 其他字符串操作

1. Fileds, Split, Join
2. Contains, Index
3. ToLower, ToUpper
4. Trim, TrimRight, TrimLeft

## 2 面向对象

### 2-1 struct

Go语言仅支持封装,不支持继承和多态

```go
// 基础结构
type treeNode struct {
    val         int
    left, right *treeNode
}

var root treeNode

root = treeNode{val: 3}
root.left = &treeNode{}
root.right = &treeNode{5, nil, nil}
root.right.left = new(treeNode)

nodes := []treeNode{
    {val: 3},
    {},
    {6, nil, nil},
}
```

工厂函数createNode

```go
func createNode(val int) *treeNode {
    // 返回局部变量的地址, 是合法的
    return &treeNode{val: val}
}

root.right.left = createNode(2)
```

为结构体提供函数方法

```go
// 为treenode提供print函数
func (node treeNode) print() {
    fmt.Print(node.val)
}
root.print()

// 传入指针来进行引用传递
func (node *treeNode) setVal(val int) {
    node.val = val
}
root.setVal(4)
```

### 2-2 指针接收者与值接收者

* **要改变内容必须使用指针接收者**
* **结构过大也要考虑使用指针接收者**
* **一致性:如有指针接收者,最好都是指针接收者**

### 2-3 封装

* 名字一般使用CamelCase
* 首字母大写:public
* 首字母小写:private

### 2-4 扩展

```go
// 1. 组合
type MyTreeNode struct {
    node *tree.Node
}

// 2. 别名
// 将slice方法进行封装
type Queue []int

func (q *Queue) Push(v int) {
    *q = append(*q, v)
}

func (q *Queue) Pop() int {
    head := (*q)[0]
    *q = (*q)[1:]
    return head
}

func (q *Queue) IsEmpty() bool {
    return len(*q) == 0
}

// 3. 内嵌(类似继承)
type MyTreeNode struct {
    *tree.Node
}

// Embedding与继承的区别:无法用父类指针指向子类
// cannot use &root (value of type *MyTreeNode) as *tree.Node value in
var baseRoot *tree.Node
baseRoot = &root
```

## 3 Build

### go mod

go mod init rep_name 初始化

go build ./... 自动在go.mod中biuld所有文件中import的依赖

go mod tidy 清理旧版本

go get 增加依赖

go get [@v...] 增加指定版本的依赖

## 4 工程化

### 4-1 接口

```go
func getRetriever() retriever {
    return infra.Retriever{}
    // return testCode.Retriever{}
}

// 接口适用于不同类型retriever
type retriever interface {
    Get(string) string
}

var r retriever = getRetriever()
```

### duck typing

download(使用者) ---> retrieve(实现者)

***接口由使用者定义***

```go
// main.go
type Retriever interface {
    // 未实现Get方法,由使用者定义
    Get(url string) string
}

// 使用者只需要使用,不需要实现
func download(r Retriever) string {
    return r.Get("http://www.imooc.com")
}

var r Retriever
r = mock.Retriever{Contents: "This is a fake imooc.com"}
r = real.Retriever{}
fmt.Println(download(r))

// mockretriever.go 具体的接口方法实现
type Retriever struct {
    Contents string
}

func (r Retriever) Get(url string) string {
    return r.Contents
}
// realretriever.go 具体的接口方法实现
type Retriever struct {
    UserAgent string
    TimeOut   time.Duration
}

func (r Retriever) Get(url string) string {
    resp, err := http.Get(url)
    if err != nil {
        panic(err)
    }

    result, err := httputil.DumpResponse(
        resp, true,
    )

    resp.Body.Close()

    if err != nil {
        panic(err)
    }

    return string(result)
}
```

### 接口的值类型

```go
/*
Retriever r:
    Type             value
    *real.Retriever  &{Mozilla/5.0 1m0s}
*/
// Type assertion: 取出r中的值
realRetriever := r.(*real.Retriever)
fmt.Println(realRetriever.UserAgent)
```

* 接口变量自带指针(指向实现者,即实现的方法)
* 接口变量同样采用值传递,几乎不需要使用接口的指针(接口中方法的实现一般是值传递实现,也可以使用指针实现)
* 指针接收者实现只能以指针方式使用;值接收者都可(如果使用指针传递实现接口中的方法,那么接口变量需要接收该方法的地址;而值传递实现的方法,两种类型都能接收) 

***interface{}可表示任意类型***

```go
type Queue []interface{}

func (q *Queue) Push(v interface{}) {
    // 限制值为int,否则运行时出错
    *q = append(*q, v.(int))  
}

// 限制传入时的值为int,否则编译出错
func (q *Queue) Push(v int) {
    *q = append(*q, v)  
}
```

### 接口的组合

组合的接口,实现者只需关注各自接口的实现方法;使用者则只需准确使用需要的接口变量中的方法(创建实现了对应方法(Get&Post)的接口变量(RetrieverPoster))

```go
// 组合Retriever与Poster
type RetrieverPoster interface {
    Retriever
    Poster
}

// 组合方法的使用
func session(s RetrieverPoster) string {
    s.Post(url, map[string]string{
        "contents": "another faked imooc.com",
    })
    return s.Get(url)
}

// 在mock中实现了Post与Get方法
s := &mock.Retriever{Contents: "This is a fake imooc.com"}
fmt.Println(session(s))
```

***特殊接口***

```go
// 相当于toString()操作,在接口中加入可以帮助格式化接口变量的输出
func (r *Retriever) String() string {
    return fmt.Sprintf(
        "Retriever: {Contents=%s}", r.Contents,
    )
}

// Reader/Writer接口
```

## 函数式编程

将func闭包进adder

```go
func adder() func(i int) int {
    sum := 0
    return func(v int) int {
        sum += v
        return sum
    }
}
```

### 函数接口

```go
// 给函数写接口
type intGen func() int

// 实现函数接口Read方法
func (g intGen) Read(p []byte) (n int, err error) {
    next := g()
    if next > 10000 {
        return 0, io.EOF
    }
    s := fmt.Sprintf("%d\n", next)
    return strings.NewReader(s).Read(p)
}
```

### 闭包的应用

二叉树的遍历

```go
func (node *Node) Traverse() {
    node.TraverseFunc(func(n *Node) {
        fmt.Printf("%d ", node.Val)
    })
    fmt.Println()
}

func (node *Node) TraverseFunc(f func(*Node)) {
    if node == nil {
        return
    }

    node.TraverseFunc(f)
    f(node)
    node.TraverseFunc(f)
}
```

统计结点数量

```go
// 统计节点数量
nodeCount := 0
root.TraverseFunc(func(node *Node) {
    nodeCount++
})
fmt.Println("Node Count: ", nodeCount)
```

通过TraverseFunc中对函数f的匿名实现来改变函数用途

### defer

defer的语句将会在程序结束时运行

```go
func tryDefer() {
    // defer将语句存入栈中
    defer fmt.Println(1)
    defer fmt.Println(2)
    fmt.Println(3)
    panic("error occurred!")
    fmt.Println(4)
}
```

### 错误处理

```go
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
```

用errWrapper进行封装,让handler实现自己的逻辑不用关心错误处理

```go
// 业务逻辑函数
type appHandler func(writer http.ResponseWriter,
    request *http.Request) error

// 针对appHandler函数,return函数做错误处理
func errWrapper(handler appHandler) func(
    http.ResponseWriter, *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        // 接收业务逻辑中的err
        err := handler(w, r)
        if err != nil {
            log.Printf("Error handling request: %s",
                err.Error())
            code := http.StatusOK
            switch {
            case os.IsNotExist(err):
                code = http.StatusNotFound
            case os.IsPermission(err):
                code = http.StatusForbidden
            default:
                code = http.StatusInternalServerError
            }
            http.Error(w,
                http.StatusText(code), code)
        }
    }
}
```

* defer + panic + recover
* Type Assertion
* 函数式编程的应用

## 并发编程

### 协程Coroutine

* 轻量级"线程"
* 非抢占式多任务处理,由协程主动交出控制权
* 编译器/解释器/虚拟机层面的多任务
* 多个协程可能在一个或多个线程上运行

```go
var a [10]int
for i := 0; i < 10; i++ {
    // 协程函数
    go func(i int) {
        for {
            // fmt.Printf("Hello from goroutine %d\n", i)
            a[i]++
        }
    }(i)
}
```

![协程调度](images/goroutine.png)

### goroutine

* 任何函数只需加上go就能送给调度器运行
* 不需要在定义时区分是否是异步函数
* 调度器在合适的点进行切换
* 使用-race来检测数据访问冲突

### channel

***Channel as first-class citizen***

```go
// 创建的channel只能发数据
func createWorker(id int) chan<- int {
    c := make(chan int)
    go func() {
        for {
            fmt.Printf("Worker %d received %c\n",
                id, <-c)
        }
    }()
    return c
}
```

***Buffered channel***

```go
c := make(chan int, 3)
go worker(0, c)
```

***Channel close and range***

channel关闭后,其中常用range来收

```go
// 1. 协程工作函数
func worker(id int, c chan int) {
    // range c来截停
    for n := range c {
        fmt.Printf("Worker %d received %d\n",
            id, n)
    }
}

c := make(chan int)
go worker(0, c)
c <- 'a'
c <- 'b'
c <- 'c'
c <- 'd'
close(c)
```

***$\color{red}{不要用共享内存来通信;通过通信来共享内存}$***

### 并发

```go
type worker struct {
    in chan int
    // done chan bool
    done func()
}
// 1. 协程工作函数
func doWork(id int, w worker) {
    for n := range w.in {
        fmt.Printf("Worker %d received %c\n",
            id, n)
        // 异步操作
        w.done()
    }
}

// 异步信号
var wg sync.WaitGroup

// 异步worker协程
wg.Add(20)
w := worker{
    in: make(chan int),
    // done: make(chan bool),
    done: func() {
        wg.Done()
    },
}
// 等待worker协程结束
wg.Wait()
```

***select***

```go
for {
    // 只能用于通道操作
    select {
    case n = <-c1:
        ...
    case n = <-c2:
        ...
    
    // 非阻塞式
    default:
        ...
    }
}
```

如果有多个 case 都可以运行,select 会随机公平地选出一个执行,其他不会执行。
否则:

* 如果有 default 子句,则执行该语句
* 如果没有 default 子句,select 将阻塞,直到某个通道可以运行；Go 不会重新对 channel 或值进行求值

***传统同步机制***

* WaitGroup
* Mutex
* Cond

### 并发编程模式

* 生成器

    ```go
    func msgGen(name string) <-chan string {
        c := make(chan string)
        go func() {
            for {
                ...
            }
        }()
        return c
    }
    ```

* 服务/任务
* 同时等待多个服务: 两种方法

    ```go
    // 接收两个channel的数据,依次输出到新的channel
    func fanIn(chs ...<-chan string) <-chan string {
        c := make(chan string)
        for _, ch := range chs {
            go func(in <-chan string) {
                for {
                    c <- <-in
                }
            }(ch) // 传入ch的复制,防止go routine冲突
        }
        return c
    }

    // select实现,仅需一个go routine
    func fanInBySelect(c1, c2 <-chan string) <-chan string {
        c := make(chan string)
        go func() {
            for {
                select {
                case m := <-c1:
                    c <- m
                case m := <-c2:
                    c <- m
                }
            }
        }()
        return c
    }
    ```

### 并发任务控制

* 非阻塞等待

    ```go
    // 非阻塞等待
    func nonBlockingWait(c <-chan string) (string, bool) {
        select {
        case m := <-c:
            return m, true
        default:
            return "", false
        }
    }
    ```

* 超时机制

    ```go
    // 超时等待
    func timeoutWait(c <-chan string, timeout time.Duration) (string, bool) {
        select {
        case m := <-c:
            return m, true
        case <-time.After(timeout):
            return "", false
        }
    }
    ```

* 任务中断/退出
* 优雅退出

    ```go
    // 主动中断任务
    func msgGenInterrup(name string, done chan struct{}) <-chan string {
        c := make(chan string)
        go func() {
            i := 0
            for {
                select {
                case <-time.After(time.Duration(rand.Intn(5000)) * time.Millisecond):
                    c <- fmt.Sprintf("service %s: message %d", name, i)
                case <-done:
                    fmt.Println("cleaning up")
                    time.Sleep(2 * time.Second) // do cleaning 
                    fmt.Println("cleaning done")
                    done <- struct{}{}  // 优雅退出:告诉main已经clearning done
                    return
                }
                i++
            }
        }()
        return c
    }
    ```
