package main

import "fmt"

//map 和silce组合

func main() {

	var s1 = make([]map[int]string,10,10)
	//没有对内部map做初始化
	s1[0] = make(map[int]string,1)
	s1[0][10] = "沙河"
	fmt.Println(s1)

	var m1 = make(map[string][]int,10)
	m1["北京"] = []int{10,20,30}
	fmt.Println(m1)
}