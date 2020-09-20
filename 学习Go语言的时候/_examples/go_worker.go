package main

import (
	"fmt"
	"time"
)

func pln(args ...interface{}) {
   fmt.Println(args...);
}

func go_worker(name string) {

	for i := 1; i <= 6; i++ {
		pln("我是一个go的协程，我的名字是", name)
		time.Sleep(1 * time.Second)
	}

	fmt.Println(name, "执行完毕！！！")
	
}

func main() {

	go go_worker("阿牛！！！")
	go go_worker("阿黄666")
   
    for i := 1; i <= 6; i++ {
		pln("我是main...")
		time.Sleep(1 * time.Second)
	}
	
}