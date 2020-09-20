package main

import (
	"fmt"
)

type People struct {
	Name string
	Age int
}

func main()  {
	peo := new(People)
	fmt.Println(peo)              //&{ 0}
	fmt.Println(peo == nil)       //false

	peo.Name = "derek"
	peo.Age = 22
	fmt.Println(peo)              //&{derek 22}

	peo2 := peo
	fmt.Println(peo2)            //&{derek 22}

	peo2.Name = "Jack"
	fmt.Println(peo, peo2)       //&{Jack 22} &{Jack 22}
}