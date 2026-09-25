package main

import(
	"fmt"
	"cse312/util"
)

func main(){
	p := util.Person(Name: "John", Age: 25)
	return p.SayHello()
}
