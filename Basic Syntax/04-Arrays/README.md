## 数组

### 声明数组

语法：
```go
    var arr_name [size]datatype
```

### 定义

```go
	// 几种数组定义初始化的方式 
	var uintArr [5]uint16 = [5]uint16{21,42,41,56,72}
	fmt.Println("uintArr=>",uintArr)

	var intArr = [5]int{-14,12,43,51,-95}
	fmt.Println("intArr=>",intArr)

	var floatArr = [...]float64{1.0,234.134,31235.21}
	fmt.Println("floatArr=>",floatArr)

	var byteArr = [...]byte{1:'a',0:'b',2:'c',5:'e',4:'d',3:'f'}
	fmt.Println("byteArr=>",byteArr)
	
	// strArr := [...]string{"str1","num","sam","str2"}
	strArr :=[...]string{0:"str1",2:"str2",1:"str3"}
	fmt.Println("strArr=>",strArr)
```

### 访问数组
```go

```

#### loop for

```go
	for i := 0 ; i < len(arr) ; i++ {
		fmt.Println(arr[i])
	}
```

#### for range

```go
	for d,c :=range arr{
		fmt.Println("下标为 %d 的数据是 %d ",d,c)
	}
```

```go
	for _,v := range arr {
		fmt.Println(v)
	}
```


### 数组切片

Go 语言切片是对数组的抽象。

Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go中提供了一种灵活，功能强悍的内置类型切片("动态数组"),与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。






