package main

import (
	"fmt"
)

func main() {
	 //n1=math.MaxFloat32 //float32最大值
	f1 := 1.234567
	fmt.Printf("%T\n",f1) //Go默认最小数都是float64类型
	f2 := float32(1.234567)
	fmt.Printf("%T\n",f2) //显示声明float32类型
	// f1 = f2					// float32类型的值不能非直接复制给float64类型的变量
}