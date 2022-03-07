package main

import "fmt"

//变量作用域

var x = 100

func f1() {
	// x := 10
	// 函数中查找变量的顺序
	// 1. 先在函数内部查找
	// 2. 找不到就往函数外面查找，一直找到全局
	fmt.Println(x)
}

func main() {
	f1()
}