package main

import (
	"fmt"
)

var (
   count = 0

)

// 递归：自己调用自己
// 递归一定要有一个明确的退出条件
// 递归适合处理那种问题，越来越小的场合
// 永远不要高估自己

//  计算nj阶乘
func f(n uint64) uint64 {
	if n <= 1{
		return 1
	}
	return n * f(n-1)
}


// 上台阶面试题
// n个台阶，一次可以走一步，也可以走两步，有多少种走法

func taijie(n uint64) uint64 {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	//if (n >= 1){
	//	println(n)
	//	count ++
	//}

	return  taijie(n-1) + taijie(n-2)
}

func main() {
	//ret := f(7)
	//fmt.Println(ret)

	//1111
	//22
	//211
	//121
	//112
	ret := taijie(3)
	fmt.Println(ret)
	//println( "count:",count)
}

