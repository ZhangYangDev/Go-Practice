## 变量 & 常量

### 声明变量 
Golang中变量的声明有三种方式：
1. 直接指定变量类型：

    语法:
    ```go
    var var_name type 
    ```
    例子：
    ```go
    var num int16 = 6792
    ```
2. 根据值类型来判断变量的类型：
   
   语法：
   ```go
   var var_name = val
   ```
   例子:
   ```go
	var tag = 'a'
	fmt.Printf("tag type:%T\n",tag)
   ```
   结果:
   ```text
   tag type:int32
   ```
3. 省略关词var来声明变量：
   语法：
    ```go
    var_name := val
    ```
   例子：
   ```go
    str := "golang"
	fmt.Printf("str type:%T\n", str)
   ```

   结果：
   ```
   str type:string
   ```

### 声明多个变量

```go
	// 声明多个变量
	
	// 声明变量 统一为指定类型
	var a ,b , c int16
	fmt.Println("a=",a,"b=",b,"c=",c)
	fmt.Printf("a type:%T\n",a)
	fmt.Printf("b type:%T\n",b)
	fmt.Printf("c type:%T\n",c)

	// 声明多个变量 然后为其赋值
	var d , e , f = 100 , "str" , 1245.8293
	fmt.Println("d=",d,"e=",e,"f=",f)
	fmt.Printf("d type:%T\n",d)
	fmt.Printf("e type:%T\n",e)
	fmt.Printf("f type:%T\n",f)

	// 省略关键字 使用:=为其赋值
	h, i, j := "var", 'l' , 983432131
	fmt.Println("h=",h,"i=",i,"j=",j)
	fmt.Printf("h type:%T\n",h)
	fmt.Printf("i type:%T\n",i)
	fmt.Printf("j type:%T\n",j)
```


PS：在Golang中声明变量之后如果没有赋初始值，编译器会使用默认值，比如 int类型默认值为0 ,String类型默认值为空字符串， 关于变量名的命名规则基本与Java无异，不再赘述


### 声明全局变量

Golang中的全局变量可以单行声明，也可以直接在`var()`中进行声明：

```go
// 声明全局变量 
var var0 byte = 'a'

var(
	var1 int16 = 5212
	var2 string 
	var3 = "variable"
)
```

## 声明常量
