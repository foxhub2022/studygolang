package main

import "fmt"

//  defer
// 第一步 返回值赋值
// 第二步 真正RET返回
// 函数中如果有defer，那么defer执行的时机是第一步和第二步之间



// defer 多余函数介绍之前释放资源（文件句柄，数据库连接，socket连接）
func deferDemo() {
	fmt.Println("start")
	defer fmt.Println("嘿嘿嘿") // defer 把它后面的语句颜值到函数即将返回的时候执行
	defer fmt.Println("呵呵呵") // 一个函数中可以有多个defer语句
	defer fmt.Println("哈哈哈") // 多个defer语句按照先进后出的顺序延迟执行
	fmt.Println("end")
}

func main() {
	deferDemo()
}
