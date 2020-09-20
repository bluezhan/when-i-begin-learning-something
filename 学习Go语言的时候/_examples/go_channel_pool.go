package main

import (
	"fmt"
	"time"
)

type Task struct {
	f func() error
}

func NewTask(arg_f func() error) *Task {
    t := Task{
    	f: arg_f,
    } 
    return &t
}

func (t *Task) Execute() {
	t.f()
}

type Pool struct {
	EntryChannel chan *Task
    JobsChannel chan *Task
    worker_num  int
}

func NewPool(cap int) *Pool {
	p := Pool{
		EntryChannel: make(chan *Task),
        JobsChannel: make(chan *Task),
        worker_num: cap,
	}

	return &p
}

func (p *Pool) worker(worker_ID int) {

	for task := range p.JobsChannel {
		task.Execute()
		fmt.Println("worker ID", worker_ID, "执行完了一个任务！")

	}

}

func (p *Pool) run() {
	for i := 0; i < p.worker_num; i++ {
		go p.worker(i)
	} 
	for task := range p.EntryChannel {
		p.JobsChannel <- task
	}

}

func main() {
	t := NewTask(func() error {
		fmt.Println(time.Now())
		return nil
	})
	p := NewPool(4)
	task_num := 0
	go func() {
        for {
        	p.EntryChannel <- t
        	task_num += 1
        	fmt.Println("当前一共执行了", task_num, "个任务！")
        }
	}()

	p.run()
}













