package main

import "fmt"

var (
	name string
	age  int
)

func main() {
	//name = "SuQiang"
	//age = 18
	//
	//fmt.Print(name)
	//fmt.Println()
	//fmt.Println(age)
	//fmt.Printf("name:%s", name)

	//for i:=1;i<10;i++{
	//	for j:=1;j<=i;j++{
	//		fmt.Printf("%d*%d=%d\t",j,i,i*j)
	//	}
	//	fmt.Println()
	//}
	//str:="hello"
	//for i,v:= range str{
	//	fmt.Printf("%d %c",i,v)
	//	fmt.Println()
	//}

	temp:=45
	if temp>18&&temp<30 {
		fmt.Println("此人已成年")
	}else if temp<10{
		fmt.Println("这是个小孩")
	}else{
		fmt.Println("这是个大叔")
	}


}
