package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {

	//s声明一个int32类型的变量x，它的值是10
	// 方法1：
	// var x int32
	// x = 10
	// 方法2：
	// var x  int32 = 10
	// 方法3 :
	// var x = int32(10)
	//  方法4 :
	// x := int32(10)
	// fmt.Println(x)



	//方法 1:
	var p  person //声明一个person类型的变量p
	p.name ="元帅"
	p.age = 18
	fmt.Println(p)
	var p1 person
	p1.name = "周林"
	p1.age = 9000
	fmt.Println(p1)
	// 方法2：
	s1 := []int{1,2,3,4}
	s2 := map[string]int{
		"stu1" : 100,
		"stu2" : 99,
		"stu3" : 0,
 	}
 	fmt.Println(s1,s2)
	// 键值对初始化
	var p2 = person{
		name : "冠华",
		age  : 15,
	}
	fmt.Println(p2)
	//值列表初始化
	var p3 = person{
		"理想",
		100,
	}
	fmt.Println(p3)

}
// 构造函数
func newPerson(name string,age int) *person {
	return &person{
		name: 	name,
		age: 	age,
	}
}