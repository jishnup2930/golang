package main

import "fmt"

var car struct{
	id int
	name string
}

type bike struct{
	id int
	name string
}

func main(){
	fmt.Println("struct")

	car.id = 1
	car.name ="bmw"
	// fmt.Println(car)

	RE := bike{
		id :1,
		name: "royal enfield",
	}
	fmt.Println(RE)

	HD := bike{
		id:2,
		name: "harley davidson",
	}
	fmt.Println(HD)

	type menuItem struct{
		name string
		price map [string]int
	}

	food := []menuItem{
		{
		name:"Biriyani",
		price : map [string]int{
			"chicken biriyani": 100,
			"beef biriyani" : 150,
		},
		},
		{
			name: "juice",
			price :map [string]int{
				"mango":100,
				"apple":110,
			},
		},
	}
	fmt.Println(food)
}