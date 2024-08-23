package main

import "log"

func main(){
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
