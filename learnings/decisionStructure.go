package main
import "log"
func main(){
	// var isTrue bool
	// isTrue = true
	// if isTrue{
		// log.Println("isTrue is ",isTrue)
	// }else{
		// log.Println("isTrue is ",isTrue)
	// }
// 
// 
    // myNym :=100
	// isTrue := false
	// if myNym > 99 && !isTrue{
	// 	log.Println("myNum is greater than 99 and isTrue is set to true")	
	// } else if myNym <100 && isTrue{
	// 	log.Println("1")
	// }else if myNym >1000 || isTrue== false{
	// 	log.Println("2")
	// }

	var myVar string
	myVar = "fish"
	switch myVar{
	case "cat":
		log.Println("cat is set to cat")
	case "dog":
		log.Println("cat is set to dog")
	case "fish":
		log.Println("cat is set to fish")
	default:
		log.Println("cat is something else")
	}


}