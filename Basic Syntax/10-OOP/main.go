package main

import (
	"fmt"
)

// 创建结构体
type User struct{
	name string
	age uint8
	gender string 
	desc string 
}


func (user *User) say() string{

	uInfo := fmt.Sprintf(user.name,user.age,user.gender,user.desc)

	return uInfo
}



func main(){
	var us = User{
		name:"Obj",
		age:30,
		gender:"male",
		desc:"sth",
	}
	fmt.Println(us.say())
}