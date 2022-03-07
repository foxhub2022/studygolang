package main


import "fmt"

type person struct {
	name 	string
	age  	int
	gender 	string
	hobby	[]string
}

func main() {
	// 声明一个person类型变量
	var p person
	p.name = "周林"
	p.age = 9000
	p.gender = "男"
	p.hobby = []string{"篮球","足球","双色球"}
	fmt.Println(p)
	fmt.Printf("%T\n",p)

	// 匿名结构体
	var s struct{
		x string
		y int
	}
	s.x = "嘿嘿嘿"
	s.y = 100
	fmt.Printf("type:%T value%v\n", s, s)
	fmt.Printf("%p\n",s)
}