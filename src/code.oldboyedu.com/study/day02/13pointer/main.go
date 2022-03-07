package main

import "fmt"

func main() {

	//n := 18
	//p := &n
	//fmt.Println(p)
	//fmt.Printf("%T\n",p)
	//m := *p
	//fmt.Println(&m)


	var a *int	//nil pointer
	fmt.Println(a)
	var a2 = new(int) //新申请一个内存地址
	fmt.Println(a2)

}
