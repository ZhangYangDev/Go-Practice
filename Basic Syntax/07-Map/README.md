## Map

### 创建Map

语法：
```go
    var_map := map[dataType]dataType{}
```

例子：
```go
	cityMap := map[string]string{}
```

### 使用Map

```go
	// map
	cityMap := map[string]string{}

	// 向Map中插入值
	cityMap [ "SZ" ] = "深圳"
    cityMap [ "BJ" ] = "北京"
    cityMap [ "SH" ] = "上海"
    cityMap [ "DJ" ] = "东京"

	// 打印Map
	fmt.Println(cityMap)

	// 遍历map中的内容
	for city := range cityMap {
		fmt.Println(city,cityMap[city])
	}

	fmt.Println("=======delete=======")

	// 删除map中的BJ
	delete(cityMap,"BJ")

	// 遍历map中的内容
	for city := range cityMap {
		fmt.Println(city,cityMap[city])
	}
	
	fmt.Println("=======update=======")
	cityMap ["SH"] = "浦西"

	// 遍历map中的内容
	for city := range cityMap {
		fmt.Println(city,cityMap[city])
	}
```

### 相关函数

```go
// 使用delete函数可以根据Map的Key来删除Value
delete() 
```


