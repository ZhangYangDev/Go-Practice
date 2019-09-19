## 接口

Inetrface类型可以定义一组方法，但是这些方法不需要实现。并且interface不能包含任何变量。然后到了某个自定义类型要使用的时候，再根据某些具体的情况吧这些方法进行实现。

接口中的所有方法都没有方法体，也就是都没有实现的方法。 在Golang中的接口，不需要进行显式的实现，只要一个变量含有接口中所有类型的方法，那么这个变量就能实现这个接口。 因此，Golang中没有 implements这样的关键词。

### 定义接口

语法：
```go
type interface_name interface{
    method1(paramsList) returnlist
    method2(paramsList) returnlist
}
```




