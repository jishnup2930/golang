package main

import {"fmt"
		"golang.org/x/exp/slices"
}

func main(){

	fmt.Println("Golang working with map")

	// var temp map[string]string
	// temp = map[string]string{
	// 	"key1":"value1",
	// }
	// fmt.Println(temp)
	
	m := map[string][]string{
		"coffee":{"cold coffee","hot coffee"},
		"tea":{"cold tea","hot tea"},
	}
	fmt.Println(m)
	m :=slices.Delete(m,1,2)

}