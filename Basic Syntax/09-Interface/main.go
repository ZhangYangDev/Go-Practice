package main 

import (
	"fmt"
)

type User interface {
	say()
}

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