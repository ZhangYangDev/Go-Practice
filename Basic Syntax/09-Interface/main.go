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

}

func (userSay UserSay)say(){
	fmt.Println("say something")
}

func main(){

	var user User

	user = new (UserSay)
	user.say()
}