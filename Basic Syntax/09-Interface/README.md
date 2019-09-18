## 接口

Inetrface类型可以定义一组方法，但是这些方法不需要实现。并且interface不能包含任何变量。然后到了某个自定义类型要使用的时候，再根据某些具体的情况吧这些方法进行实现。


### 定义接口

语法：
```go
type interface_name interface{
    method1(paramsList) returnlist
    method2(paramsList) returnlist
}
```
