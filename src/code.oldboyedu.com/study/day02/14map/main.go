package main

import "fmt"

func main() {
	var m1 map[string]int
	fmt.Println(m1 == nil) 			//还没有初始化（没有在内存中开辟空间）
	m1 = make(map[string]int,10) 	//要估算好该map容量，避免在程序运行期间再动态扩容

	m1["理想"] = 18
	m1["姬无命"] = 35


	fmt.Println(m1)
	fmt.Println(m1["理想"])
	value, ok := m1["娜扎"]
	if !ok{
		fmt.Println("查无此key")
	} else{
		fmt.Println(value)
	}

	//map遍历
	for k,v := range m1 {
		fmt.Println(k,v)
	}
	// 只遍历key
	for k := range m1 {
		fmt.Println(k)
	}

	// 只遍历value
	for _,v := range m1 {
		fmt.Println(v)
	}
}
