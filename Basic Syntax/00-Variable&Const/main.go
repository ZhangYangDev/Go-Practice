package main
import(
	"fmt"
)


// 声明全局变量 
var var0 byte = 'a'

var(
	var1 int16 = 5212
	var2 string 
	var3 = "variable"
)

// 声明全局常量
const const0 int16 = 2134
const(
	const1 byte = '1'
	const2 string = "constr"
	const3 = 2132.3232
)

func main(){
	// 全局变量

	fmt.Println("全局变量=>",var0,var1,var2,var3)
	fmt.Println("------------------------------")
	// 第一种 
	var num = 78203.5192
	fmt.Printf("num type:%T\n", num)
	fmt.Println("------------------------------")
	// 第二种
	var tag = 'a'
	fmt.Printf("tag type:%T\n",tag)
	fmt.Println("------------------------------")
	// 第三种
	str := "golang"
	fmt.Printf("str type:%T\n", str)
	fmt.Println("------------------------------")
	
	// 声明多个变量
	// 声明变量 统一为指定类型
	var a ,b , c int16
	fmt.Println("a=",a,"b=",b,"c=",c)
	fmt.Printf("a type:%T\n",a)
	fmt.Printf("b type:%T\n",b)
	fmt.Printf("c type:%T\n",c)
	fmt.Println("------------------------------")
	// 声明多个变量 然后为其赋值
	var d , e , f = 100 , "str" , 1245.8293
	fmt.Println("d=",d,"e=",e,"f=",f)
	fmt.Printf("d type:%T\n",d)
	fmt.Printf("e type:%T\n",e)
	fmt.Printf("f type:%T\n",f)
	fmt.Println("------------------------------")
	// 省略关键字 使用:=为其赋值
	h, i, j := "var", 'l' , 983432131
	fmt.Println("h=",h,"i=",i,"j=",j)
	fmt.Printf("h type:%T\n",h)
	fmt.Printf("i type:%T\n",i)
	fmt.Printf("j type:%T\n",j)
	fmt.Println("------------------------------")
	
	fmt.Println("------------------------------")
	// 全局常量
	fmt.Println("全局常量=>",const0,const1,const2,const3)
	// 声明常量
	const con1 int16 = 6728
	fmt.Println("const=> con1:",con1)
	// 可以省略数据类型标识 编译器会根据值类型进行推导
	const con2 = 7263.2344
	fmt.Printf("con2 type:%T\n",con2)

}
