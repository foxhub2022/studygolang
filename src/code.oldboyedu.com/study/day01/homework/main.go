package main

import (
	"fmt"
)

func main() {
	ret := 0
	for i:=1 ;i<10;i++ {
		//fmt.Println(i)
		//if i ==5{
		//	fmt.Printf("%d \t",i)
		//	break
		//}
		for j :=i ;j>10;j++ {
			fmt.Println("j  ")
			fmt.Printf("%d * %d = %d \t" ,i,j,j*i)
			ret = i*j
		}
	}
	fmt.Printf("结果 : %d\n",ret)
	var f1 =1.2345
	fmt.Printf("%T\n",f1)
}
