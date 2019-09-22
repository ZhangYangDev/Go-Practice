package main 

import (
	"fmt"
)
// 接口
type User interface {
	say()
}
// 结构体 
type UserSay struct {
	name string
	desc string 
}
// 接口实现
func (userSay UserSay)say(){
	fmt.Println("say something")
}

func main(){
	var user User = new (UserSay)
	// user 
	user.say()
}