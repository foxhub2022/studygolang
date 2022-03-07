package main

import "fmt"

func main() {
	//数组求和
	d := [5]int{1,3,5,7,8}
	total := 0
	for _, v := range d{
		total += v
	}
	fmt.Printf("合计：%d \n",total)

	// 计算那两个数组下标之和 等于8
	for i, v := range d{
		for j := i +1; j < len(d); j++ {
			if v + d[j] == 8 {
				fmt.Printf("这两个下标之和是8：（%v %v）\n",i,j)
			}
		}
	}
}
