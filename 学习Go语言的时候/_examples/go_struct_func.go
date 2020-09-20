package main

import "fmt"

type People struct {
	Name string
	Weight int
}

func (p People) run() {
	fmt.Println(p.Name,"正在跑步，当前体重为：",p.Weight)
}

func main()  {
	peo := People{"derek",120}
	peo.run()       //derek 正在跑步，当前体重为： 120
}