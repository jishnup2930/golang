package main

import (
	"fmt"
	"log"
)

func userProfile(){
	// for i:= 0; i <=10; i++{
		// log.Println(i)
	// }
	// animals := []string {"dog","fish","horse"}
	// for _,animals := range animals{
	// 	log.Println(animals)
	// }
	type User struct{
		firstName string
		lastName string
		email string
		age int
	}
	var users []User 
	users = append(users, User{"john","smith","johnsmith@gmail.com",25})
	users = append(users, User{"mary","jones","marujones@gmail.com",25})

	for _,l := range users{
		log.Println(l.firstName,l.lastName,l.email,l.age)
	}
	
}
func loop(){
	for i:=0; i<5; i++{
		fmt.Println(i)
	}
}

func numbers(){
	numbers:=[2] int {100,101}
	fmt.Println(numbers)
	for _,v:= range numbers{
	fmt.Println(v)}
}

func MAP(){
	letters :=map[string]string{
		"a":"A",
		"b":"B",
	}
	for key,value := range letters{
		fmt.Println(key,value)

	}
}

func main(){
	userProfile()
	loop()
	numbers()
	MAP()

}
