package main

import "fmt"

func main() {
	// 1 判断字符串汉字的数量
	// 难点是判断一个字符的汉字
	//s1 := "Hello沙河"
	//var count int
	//for _,c := range s1 {
	//	//2  判断当前这字符是不是汉字
	//	if unicode.Is(unicode.Han, c){
	//		count++
	//	}
	//}
	//// 3.把汉字出现的次数累加得到最终结果
	//fmt.Println(count)

	// how do you do 单词出现次数
	// 2.1 把字符串按空格切割得到分片
	//s2 := "how do you do"
	//s3 := strings.Split(s2, " ")
	////fmt.Println(s3)
	//// 2.2 遍历切片存储到一个map
	//m1 := make(map[string]int,10)
	//for _,w := range s3{
	//	if _, ok :=m1[w];!ok {
	//		m1[w] = 1
	//	} else {
	//		m1[w]++
	//	}
	//}
	//// 2.3 累加出现的次数
	//for key, value := range m1 {
	//	fmt.Println(key,value)
	//}


	// 回文判断
	// 字符串从左往右读和从右往左读是一样的，那么就是回文。
	// 上海自来水来自海上
	// 山西运煤车煤运西山
	// 黄山落叶松叶落山黄
	ss:= "a山西运煤车煤运西山a"
	// 解体思路：
	// 把字符串中的字符拿出来放到一个[]rune中

	r := make([]rune,0,len(ss))
	for _,c := range ss {
		r = append(r,c )
	}

	fmt.Println("[]rune:",r)
	for i :=0; i < len(r)/2; i++ {
		if r[i] != r[len(r)-1-i]{
			fmt.Println("不是回文")
			return
		}
	}
	fmt.Println("是回文")



}
