package main

import (
	"fmt"
)


func main(){
	// 函数
	
	// 函数调用
	a,b := str_split("Go","Lang")
	fmt.Println(a,b)

}

// 函数多个返回值
func str_split(x,y string)(string,string){
	return x,y
}

	