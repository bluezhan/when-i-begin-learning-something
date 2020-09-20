package main

import "fmt"

type People struct {
	Name string
	Weight int
}

func (p *People) run() {
	fmt.Println(p.Name,"正在跑步，当前体重为：",p.Weight)
	//运行一次run方法，体重减去1
	p.Weight -= 1
}

func main()  {
	peo := People{"derek",120}
	peo.run()       //derek 正在跑步，当前体重为： 120
	fmt.Println("跑完步后的体重为：",peo.Weight)        //跑完步后的体重为： 119
}