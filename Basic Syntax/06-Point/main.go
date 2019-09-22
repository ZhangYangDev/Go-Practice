package main 

import (
	"fmt"
)

func main(){

	// 声明指针
	var str string = "lang"   // 声明变量
	var sub *string			  // 声明指针变量
	fmt.Println("str的值是=>",str)
	fmt.Println("str的内存地址是=>",&str)
	
	// 将变量str的内存地址指向指针变量sub 
	sub = &str
	// 此时打印sub所存储的指针地址
	fmt.Println("sub所存储的指针地址是=>",sub)
	// 打印变量 *sub 的值
	fmt.Println("sub的值是=>",*sub)


	// 空指针 值为 nil
	var ptrFloat *float64
	fmt.Println("ptrFloat的值是=>",ptrFloat)

}