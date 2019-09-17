package main

import (
	"fmt"
)


func main(){
	// 函数
	
	// 函数调用 str_split
	a,b := str_split("Go","Lang")
	fmt.Println(a,b)

	// 函数调用 sum
	c := sum(1,2)
	fmt.Println(c)
}

// 函数多个返回值
func str_split(x,y string)(string,string){
	return x,y
}


// 求和函数
func sum(num1 , num2 int) int {
		return num1+num2
}