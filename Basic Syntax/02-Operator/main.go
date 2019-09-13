package main 

import(
	"fmt"
)

var(
	num1 uint8 = 100
	num2 uint8 = 88
)

func main(){
	// 算术运算符  自增&自减
	num1++   
	num2-- 
	fmt.Println("num1++ Val:",num1)
	fmt.Println("num2-- Val:",num2)

	// 关系运算符 大于等于&小于等于
	// num1小于等于num2吗  false
	if(num1 <= num2){
		fmt.Println("num1小于等于num2吗  true")
	}else{
		fmt.Println("num1小于等于num2吗  false")
	}
	// num1小于等于num2吗  true
	if(num1 >= num2){
		fmt.Println("num1大于等于num2吗  true")
	}else{
		fmt.Println("num1大于等于num2吗  false")
	}
		


}