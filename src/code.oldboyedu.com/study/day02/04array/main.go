package main

import "fmt"

func main() {
	var a1 [3]bool

	for _,s := range a1 {
	    fmt.Println(s)
	}

	a1 = [3]bool{true,true,true}
	fmt.Println(a1)

	// 根据初始值自动推断数组长度
	a10 := [...]int{1,2,3,4,5,6,7}
	fmt.Println(a10,len(a10))
	// 根据索引来时初始化程序
	a3 := [5]int{0:1,4:2}
	fmt.Println(a3)

	citys := [...]string{"北京","上海","深圳"}
	//数组遍历

	// 根据数据组索引遍历
	for i := 0;i < len(citys) ; i++ {
		fmt.Println(citys[i])
	}

	//for range遍历
	for i ,v := range citys{
		fmt.Println(i,v )
	}

	//多维数组
	// [[1,2] [2,3] [3,4]]
	var  a11 [3][2]int
	a11 = [3][2]int{
		[2]int{1,2},
		[2]int{2,3},
		[2]int{3,4},
	}
	fmt.Println(a11)

	//数组求和
	d := [5]int{1,3,5,7,8}
	total := 0
	for _, v := range d{
		total += v
	}
	fmt.Printf("合计：%d \n",total)

	// 计算那两个数组下标之和 等于8
	for i, v := range d{
		count1 := v
		for j := i +1; j < len(d); j++ {
			if count1 + d[j] == 8 {
				fmt.Printf("这两个下标之和是8：（%v %v）\n",i,j)
			}
		}
	}
}
