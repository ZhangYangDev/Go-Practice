package main

import "fmt"

// 变量 & 数据类型 & 常量

func main() {
	// 数值类型
	fmt.Println("--------------- 基本数据类型 Start --------------")

	// 布尔型
	var flag bool = true
	fmt.Println("bool|char=>", flag)
	// 字符串
	// string  官方将String归类为基本数据类型

	var str string = "GoLang"
	fmt.Println("string:str=>", str)
	fmt.Println("--------------- 基本数据类型 End --------------")

}

func type_number() {
	fmt.Println("数值类型")

	// 数值型  int8 int16 int32 int64  uint8 uint16 uint32 uint64
	// int 有符号  uint 无符号
	// 参考文档: https://studygolang.com/articles/7902
	var num1 int8 = -123
	fmt.Println("int8|num1=>", num1)
	var num2 uint8 = 123
	fmt.Println("unit8|num2=>", num2)
	// float 浮点型 float32 float64
	// GoLang的浮点型默认声明为64位  float64的精度要比float32的精度准确
	var fl float64 = 1.4967
	fmt.Println("float64|fl=>", fl)
}

func type_byte() {
	fmt.Println("字符型Byte===>")
	// 字符型
	// byte 可以存单个字母 类似于Java中的char
	var char byte = 'a'
	fmt.Println("byte|char=>", char)
}
