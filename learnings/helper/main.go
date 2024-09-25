package main

import (
	"fmt"
	"fun/helper"
)

 



func main(){
	fmt.Println("func demo")
	helper.PrintNames("jishnu","nandana")
	result := helper.sum(5,4)
	fmt.Println(result)
	a,b := helper.Calc(5,5)
	fmt.Println(a,b)


}