package main

import "fmt"
// 变量 & 数据类型 & 常量


func main(){

/**
Golang中变量声明有三种方式
*/
fmt.Println("--------------Start------------------")
// 第一种 声明变量不赋值 会使用默认值
var a int
fmt.Println("a=:",a)

// 第二种 根据值来判断变量的数据类型
var b = "Object"
fmt.Println("b=",b)

// 第三种 省略var 使用 := 来声明变量
c := 100
fmt.Println("c=",c)

fmt.Println("--------------End------------------")
/**
Golang中的数据类型分为基本数据类型和派生复杂数据类型
基本数据类型包含数值型、字符型、布尔型、字符串
*/



}