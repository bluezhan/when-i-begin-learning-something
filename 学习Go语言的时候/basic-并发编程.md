# 聊一聊并发编程

涉及的知识点有：函数，协程，通道（信道），锁等。


多个defer的执行顺序为“后进先出”；
defer、return、返回值三者的执行逻辑应该是：return最先执行，return负责将结果写入返回值中；接着defer开始执行一些收尾工作；最后函数携带当前返回值退出。


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

信道实例 := make(chan 信道类型)
信道实例 := make(chan 信道类型, 10)// 定义容量为10的信道

// 定义信道
pipline := make(chan int)

// 往信道中发送数据
pipline<- 200

// 从信道中取出数据，并赋值给mydata
mydata := <-pipline

信道用完了，可以对其进行关闭，避免有人一直在等待。但是你关闭信道后，接收方仍然可以从信道中取到数据，只是接收到的会永远是 0。

close(pipline)

对一个已关闭的信道再关闭，是会报错的。所以我们还要学会，如何判断一个信道是否被关闭？

当从信道中读取数据时，可以有多个返回值，其中第二个可以表示 信道是否被关闭，如果已经被关闭，ok 为 false，若还没被关闭，ok 为true。

x, ok := <-pipline

一般创建信道都是使用 make 函数，make 函数接收两个参数

第一个参数：必填，指定信道类型

第二个参数：选填，不填默认为0，指定信道的容量（可缓存多少数据）

对于信道的容量，很重要，这里要多说几点：

当容量为0时，说明信道中不能存放数据，在发送数据时，必须要求立马有人接收，否则会报错。此时的信道称之为无缓冲信道。

当容量为1时，说明信道只能缓存一个数据，若信道中已有一个数据，此时再往里发送数据，会造成程序阻塞。 利用这点可以利用信道来做锁。

当容量大于1时，信道中可以存放多个数据，可以用于多个协程之间的通信管道，共享资源。

至此我们知道，信道就是一个容器。

信道的容量，可以使用 cap 函数获取 ，而信道的长度，可以使用 len 长度获取。


按照是否可缓冲数据可分为：缓冲信道 与 无缓冲信道

缓冲信道

允许信道里存储一个或多个数据，这意味着，设置了缓冲区后，发送端和接收端可以处于异步的状态。

pipline := make(chan int, 10)
无缓冲信道

在信道里无法存储数据，这意味着，接收端必须先于发送端准备好，以确保你发送完数据后，有人立马接收数据，否则发送端就会造成阻塞，原因很简单，信道中无法存储数据。也就是说发送端和接收端是同步运行的。

pipline := make(chan int)

// 或者
pipline := make(chan int, 0)

单向信道，可以细分为 只读信道 和 只写信道。

定义只读信道

var pipline = make(chan int)
type Receiver = <-chan int // 关键代码：定义别名类型
var receiver Receiver = pipline
定义只写信道

var pipline = make(chan int)
type Sender = chan<- int  // 关键代码：定义别名类型
var sender Sender = pipline
仔细观察，区别在于 <- 符号在关键字 chan 的左边还是右边。

<-chan 表示这个信道，只能从里发出数据，对于程序来说就是只读

chan<- 表示这个信道，只能从外面接收数据，对于程序来说就是只写

遍历信道，可以使用 for 搭配 range关键字，在range时，要确保信道是处于关闭状态，否则循环会阻塞。



当信道里的数据量已经达到设定的容量时，此时再往里发送数据会阻塞整个程序。

利用这个特性，可以用当他来当程序的锁。

示例如下，详情可以看注释

package main

import {
    "fmt"
    "time"
}

// 由于 x=x+1 不是原子操作
// 所以应避免多个协程对x进行操作
// 使用容量为1的信道可以达到锁的效果
func increment(ch chan bool, x *int) {
    ch <- true
    *x = *x + 1
    <- ch
}

func main() {
    // 注意要设置容量为 1 的缓冲信道
    pipline := make(chan bool, 1)

    var x int
    for i:=0;i<1000;i++{
        go increment(pipline, &x)
    }

    // 确保所有的协程都已完成
    // 以后会介绍一种更合适的方法（Mutex），这里暂时使用sleep
    time.Sleep(3)
    fmt.Println("x 的值：", x)
}


关闭一个未初始化的 channel 会产生 panic

重复关闭同一个 channel 会产生 panic

向一个已关闭的 channel 发送消息会产生 panic

从已关闭的 channel 读取消息不会产生 panic，且能读出 channel 中还未被读取的消息，若消息均已被读取，则会读取到该类型的零值。

从已关闭的 channel 读取消息永远不会阻塞，并且会返回一个为 false 的值，用以判断该 channel 是否已关闭（x,ok := <- ch）

关闭 channel 会产生一个广播机制，所有向 channel 读取消息的 goroutine 都会收到消息

channel 在 Golang 中是一等公民，它是线程安全的，面对并发问题，应首先想到 channel。


若要程序正常执行，需要保证接收者程序在发送数据到信道前就进行阻塞状态，修改代码如下

package main

import "fmt"

func main() {
    pipline := make(chan string)
    fmt.Println(<-pipline)
    pipline <- "hello world"
}
运行的时候还是报同样的错误。问题出在哪里呢？

原来我们将发送者和接收者写在了同一协程中，虽然保证了接收者代码在发送者之前执行，但是由于前面接收者一直在等待数据 而处于阻塞状态，所以无法执行到后面的发送数据。还是一样造成了死锁。

有了前面的经验，我们将接收者代码写在另一个协程里，并保证在发送者之前执行，就像这样的代码

package main

func hello(pipline chan string)  {
    <-pipline
}

func main()  {
    pipline := make(chan string)
    go hello(pipline)
    pipline <- "hello world"
}
运行之后 ，一切正常。

包子铺里的包子已经卖完了，可还有人在排队等着买，如果不再做包子，就要告诉排队的人：不用等了，今天的包子已经卖完了，明日请早呀。

不能让人家死等呀，不跟客人说明一下，人家还以为你们店后面还在蒸包子呢。

所以这个问题，解决方法很简单，只要在发送完数据后，手动关闭信道，告诉 range 信道已经关闭，无需等待就行。

package main

import "fmt"

func main() {
    pipline := make(chan string)
    go func() {
        pipline <- "hello world"
        pipline <- "hello China"
        close(pipline)
    }()
    for data := range pipline{
        fmt.Println(data)
    }
}

说到 sync包 提供的 WaitGroup 类型。

import (
    "fmt"
    "sync"
)

func worker(x int, wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 0; i < 5; i++ {
        fmt.Printf("worker %d: %d\n", x, i)
    }
}

func main() {
    var wg sync.WaitGroup

    wg.Add(2)
    go worker(1, &wg)
    go worker(2, &wg)

    wg.Wait()
}


在 Golang 中要创建一个协程是一件无比简单的事情，你只要定义一个函数，并使用 go 关键字去执行它就行了。

如果你接触过其他语言，会发现你在使用使用线程时，为了减少线程频繁创建销毁还来的开销，通常我们会使用线程池来复用线程。

池化技术就是利用复用来提升性能的，那在 Golang 中需要协程池吗？

在 Golang 中，goroutine 是一个轻量级的线程，他的创建、调度都是在用户态进行，并不需要进入内核，这意味着创建销毁协程带来的开销是非常小的。

因此，我认为大多数情况下，开发人员是不太需要使用协程池的。

但也不排除有某些场景下是需要这样做，因为我还没有遇到就不说了。

抛开是否必要这个问题，单纯从技术的角度来看，我们可以怎样实现一个通用的协程池呢？

下面就来一起学习一下我的写法

首先定义一个协程池（Pool）结构体，包含两个属性，都是 chan 类型的。

一个是 work，用于接收 task 任务

一个是 sem，用于设置协程池大小，即可同时执行的协程数量

type Pool struct {
    work chan func()   // 任务
    sem  chan struct{} // 数量
}
然后定义一个 New 函数，用于创建一个协程池对象，有一个细节需要注意

work 是一个无缓冲通道

而 sem 是一个缓冲通道，size 大小即为协程池大小

func New(size int) *Pool {
    return &Pool{
        work: make(chan func()),
        sem:  make(chan struct{}, size),
    }
}
最后给协程池对象绑定两个函数

1、NewTask：往协程池中添加任务

当第一次调用 NewTask 添加任务的时候，由于 work 是无缓冲通道，所以会一定会走第二个 case 的分支：使用 go worker 开启一个协程。

func (p *Pool) NewTask(task func()) {
    select {
        case p.work <- task:
        case p.sem <- struct{}{}:
            go p.worker(task)
    }
}
2、worker：用于执行任务

为了能够实现协程的复用，这个使用了 for 无限循环，使这个协程在执行完任务后，也不退出，而是一直在接收新的任务。

func (p *Pool) worker(task func()) {
    defer func() { <-p.sem }()
    for {
        task()
        task = <-p.work
    }
}
这两个函数是协程池实现的关键函数，里面的逻辑很值得推敲：

1、如果设定的协程池数大于 2，此时第二次传入往 NewTask 传入task，select case 的时候，如果第一个协程还在运行中，就一定会走第二个case，重新创建一个协程执行task

2、如果传入的任务数大于设定的协程池数，并且此时所有的任务都还在运行中，那此时再调用 NewTask 传入 task ，这两个 case 都不会命中，会一直阻塞直到有任务执行完成，worker 函数里的 work 通道才能接收到新的任务，继续执行。

以上便是协程池的实现过程。

使用它也很简单，看下面的代码你就明白了

func main()  {
    pool := New(128)
    pool.NewTask(func(){
        fmt.Println("run task")
    })
}
为了让你看到效果，我设置协程池数为 2，开启四个任务，都是 sleep 2 秒后，打印当前时间。

func main()  {
    pool := New(2)

    for i := 1; i <5; i++{
        pool.NewTask(func(){
            time.Sleep(2 * time.Second)
            fmt.Println(time.Now())
        })
    }

    // 保证所有的协程都执行完毕
    time.Sleep(5 * time.Second)
}

执行结果如下，可以看到总共 4 个任务，由于协程池大小为 2，所以 4 个任务分两批执行（从打印的时间可以看出）

2020-05-24 23:18:02.014487 +0800 CST m=+2.005207182
2020-05-24 23:18:02.014524 +0800 CST m=+2.005243650
2020-05-24 23:18:04.019755 +0800 CST m=+4.010435443
2020-05-24 23:18:04.019819 +0800 CST m=+4.010499440




你吃饭吃到一半，电话来了，你一直到吃完了以后才去接，这就说明你不支持并发也不支持并行。
你吃饭吃到一半，电话来了，你停了下来接了电话，接完后继续吃饭，这说明你支持并发。
你吃饭吃到一半，电话来了，你一边打电话一边吃饭，这说明你支持并行。

并发的关键是你有处理多个任务的能力，不一定要同时。
并行的关键是你有同时处理多个任务的能力。
所以我认为它们最关键的点就是：是否是『同时』。

讲到并发，那不防先了解下什么是并发，与之相对的并行有什么区别？

这里我用两个例子来形象描述：

并发：当你在跑步时，发现鞋带松，要停下来系鞋带，这时候跑步和系鞋带就是并发状态。

并行：你跑步时，可以同时听歌，那么跑步和听歌就是并行状态，谁也不影响谁。

在计算机的世界中，一个CPU核严格来说同一时刻只能做一件事，但由于CPU的频率实在太快了，人们根本感知不到其切换的过程，所以我们在编码的时候，实际上是可以在单核机器上写多进程的程序（但你要知道这是假象），这是相对意义上的并行。

而当你的机器有多个 CPU 核时，多个进程之间才能真正的实现并行，这是绝对意义上的并行。

接着来说并发，所谓的并发，就是多个任务之间可以在同一时间段里一起执行。

但是在单核CPU里，他同一时刻只能做一件事情 ，怎么办？

谁都不能偏坦，我就先做一会 A 的活，再做一会B 的活，接着去做一会 C 的活，然后再去做一会 A 的活，就这样不断的切换着，大家都很开心，其乐融融。


















