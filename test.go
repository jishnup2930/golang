package  main
import "fmt"
func main ()  {
	fmt.Println("hi")
	var intNum int = 3215
	intNum = intNum+1
	fmt.Println(intNum)
	
	var myString  string = "It's an integer"
	fmt.Println(myString)
	fmt.Println(len(myString))

	var myBoolean bool =false
	fmt.Println(myBoolean)

	// var myMap = map[string]uint8{"Adam":54}
	// fmt.Println(myMap["Adam"])
	// age,ok =["Adam":54]
	// if ok
	// fmt.Println("Age is %v",ok)
	// name := 62
	// fmt.Println(name)

	var a int = 1
	var b int
	c :=3 
	b =2
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	var aa,bb,cc,dd int = 1,2,3,4
	fmt.Println("This is one",aa)
	fmt.Println("This is two",bb)
	fmt.Println("This is three",cc)
	fmt.Println("This is four",dd)

	var num,str =6,"Six"
	fmt.Println("This is number",num)
	fmt.Println("This is string",str)

	var(
		i int = 5
		j int  
		k string ="This is string"
	)
	fmt.Println("The out put is five",i)
	fmt.Println("The out put should be zero",j)
	fmt.Println(k)

	const A int = 5  //typed constants
	fmt.Println(A)
	const PI =3.14 //un typed constants
	fmt.Println("The PI value is ",PI)


}