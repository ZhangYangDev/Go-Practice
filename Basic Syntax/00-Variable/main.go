package main

import(
	"fmt"
)


// 声明全局变量 
var var0 byte = 'a'

var(
	var1 int16 = 5212
	var2 string 
	var3 = "variable"
)


func main(){
	// 第一种 
	var num = 78203.5192
	fmt.Printf("num type:%T\n", num)
	fmt.Println("------------------------------")
	// 第二种
	var tag = 'a'
	fmt.Printf("tag type:%T\n",tag)
	fmt.Println("------------------------------")
	// 第三种
	str := "golang"
	fmt.Printf("str type:%T\n", str)
	fmt.Println("------------------------------")
	
	// 声明多个变量
	// 声明变量 统一为指定类型
	var a ,b , c int16
	fmt.Println("a=",a,"b=",b,"c=",c)
	fmt.Printf("a type:%T\n",a)
	fmt.Printf("b type:%T\n",b)
	fmt.Printf("c type:%T\n",c)
	fmt.Println("------------------------------")
	// 声明多个变量 然后为其赋值
	var d , e , f = 100 , "str" , 1245.8293
	fmt.Println("d=",d,"e=",e,"f=",f)
	fmt.Printf("d type:%T\n",d)
	fmt.Printf("e type:%T\n",e)
	fmt.Printf("f type:%T\n",f)
	fmt.Println("------------------------------")
	// 省略关键字 使用:=为其赋值
	h, i, j := "var", 'l' , 983432131
	fmt.Println("h=",h,"i=",i,"j=",j)
	fmt.Printf("h type:%T\n",h)
	fmt.Printf("i type:%T\n",i)
	fmt.Printf("j type:%T\n",j)
	fmt.Println("------------------------------")

	
}



