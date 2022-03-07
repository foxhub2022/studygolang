package main

import (
	"fmt"
)

// 构造函数

type person struct {
	name string
	age int
}

// 构造函数：约定俗称用new开头
// 返回的是结构体还是结构体指针
// 当结构体比较大的时候，尽量使用结构体指针，减少程序的内存开销
func newPerson (name string,age int) *person {
	return  &person{
		name: name,
		age: age,
	}
}

/*

什么时候应该使用指针类型接收者

1 需要修改接受者中值
2 接受者的拷贝代价比较大对象
3 保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。
 */




// 使用值接受者 （传拷贝）
func (p person) guonian() {
	p.age++
}

// 使用指针接受者 （传对象地址）
func (p *person) zhenguonian() {
	p.age++
}

// 使用指针接受者 （传对象地址）
func (p *person) dream() {
	fmt.Println("不用上班也能挣钱")
}

func main() {
	p1 := newPerson("元帅", 18)
	//p2 := newPerson("周林", 9000)
	//p1.guonian()
	p1.zhenguonian()
	fmt.Println(p1.age)
	p1.dream()
}
