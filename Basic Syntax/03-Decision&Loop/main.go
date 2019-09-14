package main

import(
	"fmt"
)

var(
	num1 uint8 = 123
	num2 uint8 = 149
	num3 uint8 = 239
)

func main(){
	fmt.Println("==========选择语句==========")
	// 选择语句
	// if 
	if num2 > num1 {
		// 条件为True时执行
		fmt.Println("True")
	}

	// if else
	if num1 > num2 {
		fmt.Println("True")
	}else{
		fmt.Println("False")
	}

	// if else if
	if num1 > num2 {
		fmt.Println("False")
	}else if num2 > num3 {
		fmt.Println("True")	
	}else{
		fmt.Println("Other")
	}
	
	
	
	// switch  输出 123
	switch num1 {
	case 100: fmt.Println("100")
	case 119: fmt.Println("119")
	case 123: fmt.Println("123")
	}
	// switch 输出 123 
	switch {
	case num1 == 100:
		fmt.Println("100")
	case num1 == 119:
		fmt.Println("119")
	// case num1 == 123,num2 == 192:
	// 	fmt.Println("123,192")
	case num1 == 123:
		fmt.Println("123")
	}

	// select 

	fmt.Println("==========循环语句==========")
	// 循环语句	

	// for 循环
	for a := 0 ; a <= 5 ; a++ {
		fmt.Println(a)
	}
	
	// 类似于Java中的while循环 (表达式条件要为True))
	for num1 < num2 {
		num1++
		fmt.Println(num1)
	}

	// for range 
	// var arr = [5]uint8{12,134,65,81,243} 	
	// for j , k:= range  arr {
	// 	 fmt.Printf("第%d位的值 = %d\n", j,k)
	// }


	// 双重for循环 


	// 打印菱形
	
	printLozenge();
	
}



func printLozenge(){

	// 菱形上半部分  
	for z := 1 ; z <= 10 ; z++ {
		for x := 10 ; x >= z ; x-- {
			fmt.Print(" ")
		}
		for c := 1 ; c < z * 2 ; c++ {
			fmt.Print("*")
		}
		fmt.Println(" ")
	}
	// 菱形下半部分
	for v := 1 ; v <= 10 ; v++{
		
		for b := 0 ; b <= v ; b++ {
			fmt.Print(" ")
		}	
		
		for n := 19 ; n > v * 2 ; n-- {
			fmt.Print("*")
		}
		fmt.Println(" ")
	}

}