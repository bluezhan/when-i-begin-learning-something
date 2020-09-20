package main

import (
	"fmt"
	"time"
)

func pln(args ...interface{}) {
   fmt.Println(args...);
}

func worker(c chan int) {
	name := <- c
	pln(name, "执行完毕！！！")
	
}

func main() {

	c := make(chan int)

	go worker(c);

	time.Sleep(1 * time.Second)

	c <- 123456

	time.Sleep(1 * time.Second)
	
}