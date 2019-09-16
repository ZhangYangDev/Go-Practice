package main

import (
	"fmt"
)

func main(){
	
	// 声明数组
	var arr [3]uint8 
	// 给数组赋值  
	arr[0] = 213
	arr[1] = 12
	arr[2] = 172

	// 几种数组定义初始化的方式 
	var uintArr [5]uint16 = [5]uint16{21,42,41,56,72}
	fmt.Println("uintArr=>",uintArr)

	var intArr = [5]int{-14,12,43,51,-95}
	fmt.Println("intArr=>",intArr)

	var floatArr = [...]float64{1.0,234.134,31235.21}
	fmt.Println("floatArr=>",floatArr)

	var byteArr = [...]byte{1:'a',0:'b',2:'c',5:'e',4:'d',3:'f'}
	fmt.Println("byteArr=>",byteArr)
	
	// strArr := [...]string{"str1","num","sam","str2"}
	strArr :=[...]string{0:"str1",2:"str2",1:"str3"}
	fmt.Println("strArr=>",strArr)
	
	// 查看数组的默认值
	// fmt.Println("默认值=>",arr) 
	
	

	// // 循环取出数组中的数据
	// for i := 0 ; i < len(arr) ; i++ {
	// 	fmt.Println(arr[i])
	// }


	// // 使用for range来取值
	// // 第一种： 使用下标
	// for d,c :=range arr{
	// 	fmt.Println("下标为 %d 的数据是 %d ",d,c)
	// }
	// // 第二种： 不使用下标 使用_代替
	// for _,v := range arr {
	// 	fmt.Println(v)
	// }


	// 数组切片 

	// 使用append内置函数，可以对切片动态增加 

	// 留坑

}