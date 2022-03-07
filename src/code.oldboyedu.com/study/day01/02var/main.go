package main

import "fmt"

//声明变量
//var name string
//var age int
//var isOk bool

//批量声明
var (
	name string
	age int
	isOk bool
)

func main() {
	name = "理想"
	age = 16
	isOk = true
	//Go 语言中非全局变量必须要使用，不使用就编译不下去
	fmt.Println(isOk)
	fmt.Printf("name:%s",name)  //%s 占位符，使用name这个变量去替换占位符
	fmt.Println(age)   // 打印完指定的内容之后会在后面加一个换行符

	// 声明变量时，同时赋值
	var s1 string = "whb"
	fmt.Println(s1)
	// 类型推到（根据值判断该变量时什么类型）
	var s2 = 20
	fmt.Println(s2)
	//简短声明变量，只能用在函数里面
	s3 := "哈哈"
	fmt.Println(s3)
	// s1 := 10 // 同一个作用域中不能重复声明同名的变量
	// 匿名变量时一个特殊的变量: _



}