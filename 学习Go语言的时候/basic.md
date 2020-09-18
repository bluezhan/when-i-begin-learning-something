# Go语言基础笔记


## 创建变量的几种有效方式


一行声明一个变量


## 反引号的妙用：结构体里的 Tag 用法

就是对结构体数据的属性定义api


## 反射



## 数组 & 切片



切片的长度是切片中元素的数量。切片的容量是从创建切片的索引开始的底层数组中元素的数量。切片可以通过 len() 方法获取长度，可以通过 cap() 方法获取容量。数组计算 cap() 结果与 len() 相同。

new与make区别：
new只分配内存它并不初始化内存，只是将其置零。new(T)会为T类型的新项目，分配被置零的存储，并且返回它的地址，一个类型为T的值，也即其返回一个指向新分配的类型为T的指针，这个指针指向的内容的值为零（zero value），注意并不是指针为零。比如，对于bool类型，零值为false；int的零值为0；string的零值是空字符串。
make用于slice，map，和channel的初始化，返回一个初始化的(而不是置零)，类型为T的值（而不是T）。之所以有所不同，是因为这三个类型是使用前必须初始化的数据结构。例如，slice是一个三元描述符，包含一个指向数据（在数组中）的指针，长度，以及容量，在这些项被初始化之前，slice都是nil的。对于slice，map和channel，make初始化这些内部数据结构，并准备好可用的值。


Go 中数组的长度是不可改变的，而 Slice 解决的就是对不定长数组的需求。他们的区别主要有两点。

区别一：初始化方式
数组：

a := [3]int{1,2,3} //指定长度
//or
a := [...]int{1,2,3} //不指定长度
切片：

s := make([]int, 3) //指定长度
//or
s := []int{1,2,3} //不指定长度
注意 1
虽然数组在初始化时也可以不指定长度，但 Go 语言会根据数组中元素个数自动设置数组长度，并且不可改变。切片通过 append 方法增加元素：

s := []int{1,2,3} //s = {1,2,3}
s = append(s, 4) //s = {1,2,3,4}
如果将 append 用在数组上，你将会收到报错：first argument to append must be slice。

注意 2
切片不只有长度（len）的概念，同时还有容量（cap）的概念。因此切片其实还有一个指定长度和容量的初始化方式：

s := make([]int, 3, 5)
这就初始化了一个长度为3，容量为5的切片。
此外，切片还可以从一个数组中初始化（可应用于如何将数组转换成切片）：

a := [3]int{1,2,3}
s := a[:]
上述例子通过数组 a 初始化了一个切片 s。

区别二：函数传递
当切片和数组作为参数在函数（func）中传递时，数组传递的是值，而切片传递的是指针。因此当传入的切片在函数中被改变时，函数外的切片也会同时改变。相同的情况，函数外的数组则不会发生任何变化。

1. new 函数
在官方文档中，new 函数的描述如下

// The new built-in function allocates memory. The first argument is a type,
// not a value, and the value returned is a pointer to a newly
// allocated zero value of that type.
func new(Type) *Type
可以看到，new 只能传递一个参数，该参数为一个任意类型，可以是Go语言内建的类型，也可以是你自定义的类型

那么 new 函数到底做了哪些事呢：

分配内存

设置零值

返回指针（重要）

举个例子

import "fmt"

type Student struct {
   name string
   age int
}

func main() {
    // new 一个内建类型
    num := new(int)
    fmt.Println(*num) //打印零值：0

    // new 一个自定义类型
    s := new(Student)
    s.name = "wangbm"
}
2. make 函数
在官方文档中，make 函数的描述如下

//The make built-in function allocates and initializes an object //of type slice, map, or chan (only). Like new, the first argument is // a type, not a value. Unlike new, make’s return type is the same as // the type of its argument, not a pointer to it.

func make(t Type, size …IntegerType) Type

翻译一下注释内容

内建函数 make 用来为 slice，map 或 chan 类型（注意：也只能用在这三种类型上）分配内存和初始化一个对象

make 返回类型的本身而不是指针，而返回值也依赖于具体传入的类型，因为这三种类型就是引用类型，所以就没有必要返回他们的指针了

注意，因为这三种类型是引用类型，所以必须得初始化（size和cap），但是不是置为零值，这个和new是不一样的。

举几个例子

//切片
a := make([]int, 2, 10)

// 字典
b := make(map[string]int)

// 通道
c := make(chan int, 10)
3. 总结
new：为所有的类型分配内存，并初始化为零值，返回指针。

make：只能为 slice，map，chan 分配内存，并初始化，返回的是类型。

另外，目前来看 new 函数并不常用，大家更喜欢使用短语句声明的方式。

a := new(int)
*a = 1
// 等价于
a := 1
但是 make 就不一样了，它的地位无可替代，在使用slice、map以及channel的时候，还是要使用make进行初始化，然后才可以对他们进行操作。

关于 Golang 的参数传递
https://www.jianshu.com/p/21ee9cdd2df4












