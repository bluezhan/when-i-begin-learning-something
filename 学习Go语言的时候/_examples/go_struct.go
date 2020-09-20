package main

import (
	"fmt"
)

type People struct {
	Name string
	Age int
}

func main()  {
	var peo People
	fmt.Println(peo)           //{ 0}
	fmt.Printf("%p", &peo)     //0x110040f0

	//赋值
	//第一种
	peo = People{"derek", 20}
	fmt.Println(peo)      //{derek 20}
	//第二种
	peo2 := People{Age: 12, Name: "jack"}
	fmt.Println(peo2)     //{jack 12}

	peo3 := &People{Age: 18, Name: "tom"} // new(People)
	fmt.Println(peo3)     //&{tom 18}

	//第三种
	peo.Name = "alice"
	peo.Age = 25
	fmt.Println(peo)     //{alice 25}


}