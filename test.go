package  main
import "fmt"
func one ()  {
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
	time:=20
	if time<18{
		fmt.Println("Time is less than 18")
		
	}else{
		fmt.Println("Time is greater than 18")
	}

	if time<10{
		fmt.Println("Morning")
	}else if time<20{
		fmt.Println("After noon")
	}else {
		fmt.Println("Evening")
	}

	// var number int

	// fmt.Println("Enter a number between 1 and 5")
	// fmt.Scanln(&number)
	// switch number {
	// case 1:
	// 	fmt.Println("The number is one")
	// case 2:
	// 	fmt.Println("The number is two")
	// case 3:
	// 	fmt.Println("The number is three")
	// case 4:
	// 	fmt.Println("The number is four")
	// case 5:
	// 	fmt.Println("The number is five")
	// default:
	// 	fmt.Println("Number is out of range")
	// }
	// switch number{
	// case 1,3,5:
	// 	fmt.Println("The number is odd")
	// case 2,4:
	// 	fmt.Println("The number is even")
	// }
	for i:=1;i<=5;i++{
		fmt.Println(i)
	}
	var adj =[2]string{"Tasty","Big"}
	var fruits =[3]string{"apple","orange","banana"}
	for i:=0; i<len(adj); i++{
		for j:=0; j<len(fruits); j++{
			fmt.Println(adj[i],fruits[j])
		}
	// fruits := [3]string{"apple", "orange", "banana"}
	for val:= range fruits {
		fmt.Printf("%v\n", val)
	}
	}

	}	

func two(){
	fmt.Println("The function one is executed")
}
func name(fname string){
	fmt.Println("Hello",fname)
}

func three(fname string , age int){
	fmt.Println("Name:",fname,"Age:",age)
}
func add(x int,y int)(result int){
	result=x + y
	return
}
func multipleReturn(x int,y string)(result int,text string){
	result = x*x
	text = y+"World"
	return

}

func main(){
	one()
	two()
	fname:=[2]string{"jishnu","nandana"}
	for i:=0;i<len(fname);i++{
		name(fname[i])
	}
	three("Jishnu",27)
	fmt.Println(add(45,5))
	fmt.Println(multipleReturn(4,"Hello"))
}