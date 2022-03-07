package main

import "fmt"

// 切片 slice
func main() {

	//切片定义 切片是一个引用类型，它的内部结构包含地址、长度和容量。切片一般用于快速地操作一块数据集合。
	var s1 []int 	//定义一个存放int类型的切片
	var s2 []string //定义一个存放string类型的切片
	fmt.Println(s1,s2)

	// 初始化
	s1 = []int{1,2,3}
	s2 = []string{"沙河","张江","平山村"}
	fmt.Println(s1,s2)
	//长度和容量
	fmt.Println("len(s1):%d cap(s1):%d",len(s1),cap(s1))
	fmt.Println("len(s1):%d cap(s1):%d",len(s2),cap(s2))

	var a1 []int = []int{1,2,3}
	a2 := a1
	a2[0] = 100
	fmt.Println(a1,a2)

	fmt.Println("**********************************************************")


	c1 := []int{1,3,5,7,9,11,13}
	s3 := c1[0:4] //基于一个数组切割，左包含右不包含（左闭右开）
	fmt.Println(s3)
	s4 := c1[1:6]
	fmt.Println(s4)
	s5 := c1[:4] //[0:4] [1 3 5 7]
	s6 := c1[3:] //[3:len(c1)] [ 7 9 11 13]
	s7 := c1[:]  // [0:len(c1)]
	fmt.Println(s5,s6,s7)
	//切片容量是指底层数据组容量
	fmt.Printf("len(s5): %d  cap(s5): %d \n",len(s5),cap(s5))
	// 底层数组从切片的第一个元素到最后的元素数量
	fmt.Printf("len(s6): %d  cap(s6): %d \n",len(s6),cap(s6))
	//切片再切割
	s8 := s6[:3] //[13]
	fmt.Printf("len(s8): %d  cap(s8): %d \n",len(s8),cap(s8))
	s6[2] = 100
	fmt.Println(c1)
	fmt.Println(s8)






}
