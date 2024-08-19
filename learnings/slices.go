package main

import (
	"log"
)

func main(){
	var mySlice[] string
	mySlice = append(mySlice,"jishnu") 
	mySlice = append(mySlice, "nandana")

	var myInt [] int
	myInt = append(myInt, 8)
	myInt = append(myInt, 4)
	

	log.Println(mySlice)
	log.Println(myInt)

	names :=[]string {"one","seven"}
	log.Println(names)

}