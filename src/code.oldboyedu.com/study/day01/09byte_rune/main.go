package main

import (
	"fmt"
	"unicode"
)

//abcd 可以成为byte
// 中文字符 定义成 rune
// rune 类型代表一个UTF-8字符

// Go 语言中为了处理非ASCII码类型的字符 定义了新的rune类型
func main() {
	//s := "Hello 沙河"
	//n := len(s)
	//fmt.Println(n)


	//for i := 0; i<len(s);i++ {
	//	fmt.Println(s[i])
	//	fmt.Printf("%c\n",s[i]) //%c: 字符
	//}

	//for _, c :=range s {
	//	fmt.Printf("%c\n",c)
	//}

	//字符串修改
	//s2 := "白萝卜"
	//s3 := []rune(s2) //把字符串强制转换成一个rune切片
	//s3[0] = '红'
	//fmt.Println(string(s3)) //把rune切片强制转成字符串

	//c1 := "红"	//string
	//c2 := '红'	// rune(int32)
	//fmt.Printf("c1 %T c2 %T \n",c1,c2)
	//c3 := "H"  		// string
	//c4 := byte('H') //byte(uint8)
	//fmt.Printf("c3 %T c4 %T \n",c3,c4)

	// 类型转换
	//n1 := 10 //int
	//var f float64
	//f = float64(n1)
	//fmt.Println(f)
	//fmt.Printf("%T \n",f)





	//计算中文字符个数
	word := "Hello沙河小王子"
	word1 := []rune(word)
	//fmt.Println(len(word))
	//fmt.Printf("字符长度 %i",len(word))
	count := 0
	for _, d := range word1 {
		fmt.Println(string(d))

		if unicode.Is(unicode.Han,d){
			count ++
		}
		//fmt.Println(d>100)
		//fmt.Println(_)
		//fmt.Printf("次数: %d",count)
		//count ++
	}
	fmt.Printf("有 %v 个汉字",count)


}
