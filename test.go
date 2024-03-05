package  main
import "fmt"
func main ()  {
	// fmt.Println("hi")
	// var intNum int = 3215
	// intNum = intNum+1
	// fmt.Println(intNum)
	
	// var myString  string = "It's an integer"
	// fmt.Println(myString)
	// fmt.Println(len(myString))

	// var myBoolean bool =false
	// fmt.Println(myBoolean)

	// var myMap = map[string]uint8{"Adam":54}
	// fmt.Println(myMap["Adam"])
	// age,ok =["Adam":54]
	// if ok
	// fmt.Println("Age is %v",ok)
	// name := 62
	// fmt.Println(name)

	// var a int = 1
	// var b int
	// c :=3 
	// b =2
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)

	// var aa,bb,cc,dd int = 1,2,3,4
	// fmt.Println("This is one",aa)
	// fmt.Println("This is two",bb)
	// fmt.Println("This is three",cc)
	// fmt.Println("This is four",dd)

	// var num,str =6,"Six"
	// fmt.Println("This is number",num)
	// fmt.Println("This is string",str)

	// var(
	// 	i int = 5
	// 	j int  
	// 	k string ="This is string"
	// )
	// fmt.Println("The out put is five",i)
	// fmt.Println("The out put should be zero",j)
	// fmt.Println(k)

	// const A int = 5  //typed constants
	// fmt.Println(A)
	// const PI =3.14 //un typed constants
	// fmt.Println("The PI value is ",PI)

	// const (
	// 	B = 1
	// 	C int=6
	// 	D = "jishnu"
	// )
	// 	fmt.Printf("%d\n%d\n%s\n",B,C,D)

	// var I = 7
	// var txt ="text"
	// fmt.Printf("%v\n",I)
	// fmt.Printf("%#v\n",i)
	// fmt.Printf("%T\n",I)

	// fmt.Printf("%v\n",txt)
	// fmt.Printf("%#v\n",txt)
	// fmt.Printf("%T\n",txt)

	// var ab int = 5
	// var myText = "Text"
	// var boo bool = true
	// var fl float32 = 54444.45
	// fmt.Println(ab)
	// fmt.Println(myText)
	// fmt.Println(boo)
	// fmt.Println(fl)

	// var x float64 =1.6e+308
	// fmt.Printf("Type: %T,value:%v \n",x,x)

	// var arr =[3]int{1,2,3}
	// car := [...]string{"BMW","AUDI","Innova","Porshe"}
	// fmt.Println(arr)
	// fmt.Println(car)
	// fmt.Println(car[2])

	// car[2] ="Ferrari"
	// fmt.Println(car)

	// fmt.Println(len(car))

	myslice := []int{1,2,3,4}
	fmt.Println(len(myslice))
	fmt.Println(cap(myslice))
	fmt.Println(myslice)

	myslice1 := []string{"a","b","c"}
	fmt.Println(len(myslice1))
	fmt.Println(cap(myslice1))
	fmt.Println(myslice1)

	slice:=[5]int{1,2,3,4,5}
	arr:=slice[1:4]
	fmt.Printf("slice=%v\n",arr)
	fmt.Printf("length = %d\n", len(arr))
	fmt.Printf("capacity = %d\n", cap(arr))

	slice[2]=6
	fmt.Println(slice)
	slice2:=append(myslice,10,20,30)
	fmt.Println(slice2)

	//operators
	var a = 10+10
	var sum = a +20
	var sum1 = sum+a
	fmt.Println(sum1)

	fmt.Println("Comparison operator",a>sum)
	fmt.Println("Comparison operator",a<sum)
	fmt.Println(a==sum)


	if sum>a{
	fmt.Println("Sum is greater than a ")
	}
	x:=10
	y:=39
	if x<y{
		fmt.Println("x is smaller than y")
	}
}	