package main

import "fmt"


// 给自定义类型加方法
// 不能给别的包里面类型加防范，只能给自己包的类型加方法


type myInt int

func (m myInt) hello() {
	fmt.Println("我是一个Int")
}

func main() {
	m := myInt(100)
	m.hello()
	println(m)
}
