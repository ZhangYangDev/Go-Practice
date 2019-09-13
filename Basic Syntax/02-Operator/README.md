## 运算符

```go
    var(
        num1 uint8 = 100
        num2 uint8 = 88
    )

```

### 算数运算符
```go
    --: 递减
    ++: 递增
```

```go
	num1++   
	num2-- 
	fmt.Println("num1++ Val:",num1)
	fmt.Println("num2-- Val:",num2)
```



```
    num1++ Val: 101
    num2-- Val: 87
    num1小于等于num2吗  false
    num1大于等于num2吗  false
```
PS: Golang不支持三元表达式可以使用 `if else`来替代。
