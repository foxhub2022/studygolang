package main

import "fmt"

type dog struct {
	name string
}

//构造函数
func newDog( name string) dog {
	return dog {
		name: name,
	}
}

//func (p )

// 方法是作用于特定类型的函数
// 接受者表示的事调用该方法的具体变量，多用类型名称的小写字符
func (d dog) wang(){
	fmt.Printf("%s 汪汪汪", d.name)
}

func main() {
	d1 := newDog("zhoulin")
	d1.wang()
}