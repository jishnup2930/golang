package main

import ("fmt"
        "time"
)

var s = "seven"


func saySomething()(string,string){
	return "something","else"
}

// package main
// import ("fmt")

// func main() {
//   var student1 string = "John" //type is string
//   student2 := "Jane" //type is inferred
//   x := 2 //type is inferred

//   fmt.Println(student1)
//   fmt.Println(student2)
//   fmt.Println(x)
// }

type UserDetails struct{
	firstName   string
	lastName    string
	phoneNumber string
	age         int
	dateOfBirth time.Time
}
func user(){
	// var name UserDetails
	// UserDetails={
	// 	firstName:"JISHNU",
	// 	lastName :"NANDANA",
	// 	phoneNumber:"9595987856",
	// 	age :27,
	// 	// dateOfBirth:12/01/1997,
	// }
	// fmt.Println(User)
	name2 :=UserDetails{
		lastName: "nandana",
	}
	var name UserDetails
	name.firstName="jishnu"
}
func (m *UserDetails) printFirstName()string{
	// fmt.Println("My first name is ",m.firstName)
	return m.firstName

}
func main(){
	// fmt.Println("Hello world")
	// var whatToSay string ="nanadana"
	// var i int =45
	// // whatToSay = "jishnu"
	// fmt.Println(whatToSay)
	// fmt.Println(i)
	// one,two :=saySomething()
	// fmt.Println(one,two) 
	// type name  struct{
	// 	a string
	// 	b string
	// }
	// n:=name{
	// 	a:"jishnu",
	// 	b:"nandana",
	// }
	// fmt.Println(n)
	fmt.Println("My first name is ",name.printFirstName())
}