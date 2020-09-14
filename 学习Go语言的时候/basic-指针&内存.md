# 让我们聊一聊指针


首先说结论：在Go语言里，所有的参数传递都是值传递(传值)，都是一个副本，一个拷贝，因为拷贝的内容有时候是非引用类型（int、string、struct等这些），这样就在函数中就无法修改原内容数据；有的是引用类型（指针、map、slice、chan等这些），这样就可以修改原内容数据。

非引用类型（值类型）：int，float，bool，string，以及数组和struct
特点：变量直接存储值，内存通常在栈中分配，栈在函数调用完会被释放

引用类型：指针，slice，map，chan，接口，函数等
特点：变量存储的是一个地址，这个地址存储最终的值。内存通常在堆上分配，当没有任何变量引用这个地址时，该地址对应的数据空间就成为一个垃圾，通过GC回收

// modify s[0] value
func modify(s1 []int) {
    s1[0] += 100
}

func main() {
    a := [5]int{1, 2, 3, 4, 5}
    s := a[:]
    modify(s)
    fmt.Println(s[0])
}

// Output:
// 101

---

// modify s[0] value
func modify2(s []int) {
    fmt.Printf("%p \n", &s)
    fmt.Printf("%p \n", &s[0])
    s[0] += 100
    fmt.Printf("%p \n", &s)
    fmt.Printf("%p \n", &s[0])
}

func main() {
    a := [5]int{1, 2, 3, 4, 5}
    s := a[:]

    fmt.Printf("%p \n", &s)
    fmt.Printf("%p \n", &s[0])
    modify2(s)
    fmt.Println(s[0])
}

// Output:
// 0xc04203c400 
// 0xc042039f50 
// 0xc04203c440 
// 0xc042039f50 
// 0xc04203c440 
// 0xc042039f50 
// 101

--- 

// modify s
func modify(s []int) {
    fmt.Printf("%p \n", &s)
    s = []int{1,1,1,1}
    fmt.Println(s)
    fmt.Printf("%p \n", &s)
}

func main() {
    a := [5]int{1, 2, 3, 4, 5}
    s := a[:]
    fmt.Printf("%p \n", &s)
    modify(s)
    fmt.Println(s[3])
}

// Output:
// 0xc042002680 
// 0xc0420026c0 
// [1 1 1 1]
// 0xc0420026c0 
// 4

---

func modifyElem(a int) {
    a += 100
}

func modifyArray(a [5]int) {
    a = [5]int{5,5,5,5,5}
}

func main() {
    var s = [5]int{1, 2, 3, 4, 5}
    modifyElem(s[0])
    fmt.Println(s[0])
    modifyArray(s)
    fmt.Println(s)
}

// Output:
// 1
// [1 2 3 4 5]

---

var i int = 5

func main() {
    modify(&i)
    fmt.Println(i)
}

func modify(i *int) {
    *i = 6
}

// Output:
// 6


说一下内存

我们在编程的时候，实际上就是在操作内存，除非是进行IO操作写磁盘。其余的不管你是一半的变量还是Hibernate的Entity，都是在内存中闪转腾挪。

我上学的时候，C语言课程是第一门编程语言课程，其中最难的部分就是指针，而指针就是直接操作内存的，所谓的C语言是最接近底层的语言，其中很重要的原因就是以为C语言让程序员可以直接去动内存。

其实在很多年前，人们编程的时候绝对不像想在这么幸福，总是要直接操作内存的，而更久远一点的程序员们，要用汇编语言直接写指令，再久一点的程序员，就要在纸带上打孔，用01010这种二进制编码编程了。

我说了这么多，其实想说的是，现在的很多编程语言比如Java，其实是对程序员隐藏了其内存操作的细节的。

我们都知道Java有堆内存和栈内存，堆内存里是实际的对象，栈内存中的变量指向了对象，这里的指向，其实就是指针了。那么指向的是什么？有没有人曾经思考过这个问题，在内存中，如何快速的寻找一个值？

答案自然是地址，只有用地址访问是最快的。


https://www.jianshu.com/p/61b1a958e3ab

指针
指针是存储变量内存地址的变量，表达了这个变量在内存存储的位置。

我们常说：指针指向了变量。

Go中指针变量类型为*T，指向一个T类型的变量。

通过&操作符用于获取一个变量的地址。

b := 255 //声明变量
var a *int = &b // 通过&b获取变量地址，指针变量a指向了b，a值是b的地址
获取指针指向变量的值
&b是获取变量b的地址。
a = &b，将a指针指向b。
通过*a获取指针指向变量的值，我们叫反解引用。
函数使用传递指针参数
我们经常有类似的需求：

func change(val *int){
    *val = 55
}
看个复杂的例子：

a := 58 //变量a赋值 a=58
b := &a // b指向a，b存储了变量a的地址
change(b) //方法接受一个指针变量，在方法中进行反解，并赋值为55
// a此时被修改为55
我们可以把一个数组通过指针传递传入一个方法，但是理解起来比较复杂，不是Go推荐的写法，这种情况通常采用切片来处理。

func modify(sis []int){
    sls[0] = 90
}
 
func main(){
    a := [3]int{89,90,91}
    modify(a[:])
    // 输出90,90,91
}
所以这种场景还是传递切片吧。


https://www.jianshu.com/p/61b1a958e3ab


