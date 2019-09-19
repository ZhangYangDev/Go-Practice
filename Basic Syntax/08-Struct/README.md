## 结构体

Golang中没有类(class)的概念，结构体(stuct)和其它语言中的类有着同等的地位，可以认为Golang是基于struct结构体来实现OOP的特性的


### 声明结构体

语法：
```go
type struct_name struct{
    var_name type
    ......
}
```
例子：
```go
type User struct {
	name string
	gender string 
	age uint8
	desc string 
}
```
