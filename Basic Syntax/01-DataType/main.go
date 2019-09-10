package main

import(
	"fmt"
	"reflect"
)

// 变量 & 数据类型 & 常量

func main() {
	// 数值类型
	fmt.Println("--------------- 数值类型 --------------")

	// int 
	var num1 int8 = -82
	var num2 uint16= 32834

	fmt.Printf("num1 type:", reflect.TypeOf(num1))
	fmt.Printf("\nnum2 type:", reflect.TypeOf(num2))
	// float
	var num3 float64 = 231.4356
	var num4 complex128 = 23434.4324
	fmt.Printf("\nnum3 type:", reflect.TypeOf(num3))	
	fmt.Printf("\nnum4 type:", reflect.TypeOf(num4))	

	// rune  类似于int32
	var num5 rune = 554
	fmt.Println("\nnum5 type:", reflect.TypeOf(num5))
	
	// 布尔型
	var flag bool = true
	fmt.Println("\n--------------- 布尔型 --------------")
	fmt.Println("\nflag type:", reflect.TypeOf(flag))
	
	// 字符串
	// string  官方将String归类为基本数据类型
	var str = "GoLang"
	fmt.Println("\n--------------- 字符串 --------------")
	fmt.Printf("\nstr type:", reflect.TypeOf(str))

	// 字符型 
	fmt.Println("\n--------------- 字符型 --------------")
	var bit byte = 'm'
	fmt.Printf("\nbit type:", reflect.TypeOf(bit))

}