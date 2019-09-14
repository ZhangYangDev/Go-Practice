## 选择语句

在Golang中,if语句的条件部分是可以省略掉括号的,除此之外基本与Java中的if使用没有区别

```go
    var(
        num1 uint8 = 123
        num2 uint8 = 149
        num3 uint8 = 239
    )
```

### if

```go
    // 示例：
	if num2 > num1 {
		// 条件为True时执行
		fmt.Println("True")
    }
    // 结果： True

```


### if else 

```go
    // 示例：
	if num1 > num2 {
		fmt.Println("True")
	}else{
		fmt.Println("False")
    }
    // 结果：False
```

### if else if 

```go
    // 示例：
    if num1 > num2 {
		fmt.Println("False")
	}else if num2 > num3 {
		fmt.Println("True")	
	}else{
		fmt.Println("Other")
    }
    // 结果：Other
```

### switch

#### 第一种写法

```go
	switch num1 {
	case 100: fmt.Println("100")
	case 119: fmt.Println("119")
	case 123: fmt.Println("123")
	}

```
#### 第二种写法
```go
	switch {
	case num1 == 100:
		fmt.Println("100")
	case num1 == 119:
		fmt.Println("119")
	// case num1 == 123,num2 == 192:
	// 	fmt.Println("123,192")
	case num1 == 123:
		fmt.Println("123")
	}
```

## 循环语句


### for
```go
	for a := 0 ; a <= 5 ; a++ {
		fmt.Println(a)
	}
```



### 使用for循环打印出一个菱形 
```go
	// 菱形上半部分  
	for z := 1 ; z <= 10 ; z++ {
		for x := 10 ; x >= z ; x-- {
			fmt.Print(" ")
		}
		for c := 1 ; c < z * 2 ; c++ {
			fmt.Print("*")
		}
		fmt.Println(" ")
	}
	// 菱形下半部分
	for v := 1 ; v <= 10 ; v++{
		
		for b := 0 ; b <= v ; b++ {
			fmt.Print(" ")
		}	
		
		for n := 19 ; n > v * 2 ; n-- {
			fmt.Print("*")
		}
		fmt.Println(" ")
	}
```
