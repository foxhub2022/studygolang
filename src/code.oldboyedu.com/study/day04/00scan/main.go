package main

import (
"fmt"
)

func main() {
	var (
		name string
		age  int
	)
	fmt.Print("输入姓名和年龄，使用空格分隔：")
	fmt.Scanln(&name, &age)
	fmt.Printf("name: %s\nage: %d\n", name, age)
}