package main

import (
	"fmt"
)

func main() {
	//var name string
	//name = "理想"
	//fmt.Println(name)
	//var ages [30]int
	//ages = [30]int{1,2,3,4}
	//fmt.Println(ages)
	//var ages2 = [...]int{1,2,3,4,5,6,7}
	//fmt.Println(ages2)
	//var ages3 = [...]int{1:100,20:200}
	//fmt.Println(ages3)

	// 切片
	//var s1 []int
	//fmt.Println(s1)
	//fmt.Println(s1 == nil)
	//s1 = []int{1,2,3}
	//fmt.Println(s1)
	//// make初始化
	//s2 := make([]bool,2,4)
	//fmt.Println(s2)
	//s3 := make([]int,0,4)
	//fmt.Println(s3 == nil)


	//var s1 []int
	//s1 = append(s1,1)
	//fmt.Println(s1)

	//s1 := []int{1,2,3}
	//s2 := s1
	//var s3 = make([]int,2,3)
	//copy(s3,s1)
	//fmt.Println(s2) // [1 2 3]
	//s2[1] = 100
	//fmt.Println(s2) // [1 200 3 ]
	//fmt.Println(s1) // [1 200 3 ]
	//fmt.Println(s3)

	// 指针
	// GO只能读取不能修改，
	addr := "沙河"
	addrP := &addr
	fmt.Println(addrP) //内存地址
	fmt.Printf("%T\n",addrP)


	//map
	var m1 map[string]int
	fmt.Println(m1 == nil)
	m1 = make(map[string]int,10)
	m1["理想"] = 100
	fmt.Println(m1)
	fmt.Println(m1["ji"]) //如果key不存在，返回值是value对应类型
	// 如果返回值是布尔型，我们通常用ok去接收
	score, ok := m1["ji"]
	if !ok {
		fmt.Println("没有姬无命这个人")
	} else {
		fmt.Println("姬无命的分数是",score)
	}
	delete(m1,"lixiang") //删除key不存在什么也不干
	delete(m1,"理想")
	fmt.Println(m1)
	fmt.Println(m1 == nil)


}
