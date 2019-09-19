package main


import (
	"fmt"
)

// 定义结构体
type User struct {
	name string
	gender string 
	age uint8
	desc string 
}


type StructTest struct{

}

func main(){
	fmt.Println(User{name:"Sam",gender:"male",age:27,desc:"I'm a Man"})
}