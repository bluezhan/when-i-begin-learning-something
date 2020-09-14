# 聊一聊并发编程

涉及的知识点有：函数，协程，通道（信道），锁等。


golang里面有两个保留的函数：init函数（能够应用于所有的package）和main函数（只能应用于package main）。这两个函数在定义时不能有任何的参数和返回值。

虽然一个package里面可以写任意多个init函数，但这无论是对于可读性还是以后的可维护性来说，我们都强烈建议用户在一个package中每个文件只写一个init函数。

go程序会自动调用init()和main()，所以你不需要在任何地方调用这两个函数。每个package中的init函数都是可选的，但package main就必须包含一个main函数。

由于 Go语言是编译型语言，所以函数编写的顺序是无关紧要的，它不像 Python 那样，函数在位置上需要定义在调用之前。

func sum(a int, b int) (int){
    return a + b
}
func main() {
    fmt.Println(sum(1,2))
}

使用 ...int，表示一个元素为int类型的切片，用来接收调用者传入的参数。

// 使用 ...类型，表示一个元素为int类型的切片
func sum(args ...int) int {
    var sum int
    for _, v := range args {
        sum += v
    }
    return sum
}
func main() {
    fmt.Println(sum(1, 2, 3))
}

// output: 6

其中 ... 是 Go 语言为了方便程序员写代码而实现的语法糖，如果该函数下会多个类型的函数，这个语法糖必须得是最后一个参数。

同时这个语法糖，只能在定义函数时使用。

上面那个例子中，我们的参数类型都是 int，如果你希望传多个参数且这些参数的类型都不一样，可以指定类型为 ...interface{} （你可能会问 interface{} 是什么类型，它是空接口，也是一个很重要的知识点，可以这篇文章查看相关内容），然后再遍历。

比如下面这段代码，是Go语言标准库中 fmt.Printf() 的函数原型：

import "fmt"
func MyPrintf(args ...interface{}) {
    for _, arg := range args {
        switch arg.(type) {
            case int:
                fmt.Println(arg, "is an int value.")
            case string:
                fmt.Println(arg, "is a string value.")
            case int64:
                fmt.Println(arg, "is an int64 value.")
            default:
                fmt.Println(arg, "is an unknown type.")
        }
    }
}

func main() {
    var v1 int = 1
    var v2 int64 = 234
    var v3 string = "hello"
    var v4 float32 = 1.234
    MyPrintf(v1, v2, v3, v4)
}
在某些情况下，我们需要定义一个参数个数可变的函数，具体传入几个参数，由调用者自己决定，但不管传入几个参数，函数都能够处理。


上面提到了可以使用 ... 来接收多个参数，除此之外，它还有一个用法，就是用来解序列，将函数的可变参数（一个切片）一个一个取出来，传递给另一个可变参数的函数，而不是传递可变参数变量本身。

同样这个用法，也只能在给函数传递参数里使用。

例子如下：

import "fmt"

func sum(args ...int) int {
    var result int
    for _, v := range args {
        result += v
    }
    return result
}

func Sum(args ...int) int {
    // 利用 ... 来解序列
    result := sum(args...)
    return result
}
func main() {
    fmt.Println(Sum(1, 2, 3))
}


Go支持返回带有变量名的值，这个返回值写法很怪

func double(a int) (b int) {
    // 不能使用 := ,因为在返回值哪里已经声明了为int
 b = a * 2
    // 不需要指明写回哪个变量,在返回值类型那里已经指定了
 return
}
func main() {
 fmt.Println(double(2))
}
// output: 4



回调函数使用

// 第二个参数为函数
func visit(list []int, f func(int)) {
    for _, v := range list {
        // 执行回调函数
        f(v)
    }
}
func main() {
    // 使用匿名函数直接做为参数
    visit([]int{1, 2, 3, 4}, func(v int) {
        fmt.Println(v)
    })
}

一个 goroutine 本身就是一个函数，当你直接调用时，它就是一个普通函数，如果你在调用前加一个关键字 go ，那你就开启了一个 goroutine。

// 执行一个函数
func()

// 开启一个协程执行这个函数
go func()

一个 Go 程序的入口通常是 main 函数,程序启动后，main 函数最先运行，我们称之为 main goroutine。

在 main 中或者其下调用的代码中才可以使用 go + func() 的方法来启动协程。

main 的地位相当于主线程，当 main 函数执行完成后，这个线程也就终结了，其下的运行着的所有协程也不管代码是不是还在跑，也得乖乖退出。










