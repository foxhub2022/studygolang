package main

import "fmt"

func adder() func(int) int {
	var x = 100
	return func(y int) int {
		x += y
		return x
	}
}

func f1(f func())  {
	fmt.Println("this is f1")
	f()

}
func f2(x, y int)  {
	fmt.Println("this is f2")
	fmt.Println(x + y)
}
//要求：
// f1(f2)

func f3(f func(int,int),x,y int) func() {
	tmp := func ()    {
      f(x,y)
    }
    return tmp
}

func main() {

	//ret := adder()
	//fmt.Printf("%T\n",ret)
	//ret2 := ret(200)
	//fmt.Println(ret2)

	//f1( f2(1,2) )
	ret := f3(f2,100,200)
	fmt.Printf("%T\n",ret)
	//ret()

	f1(ret)


}

