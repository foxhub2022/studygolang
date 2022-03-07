package main

const (
	c1 = 100  //100
	c2		  //100
	c3		  //100
)


// iota
const (

	a1 = iota //0
	a2        //1
	a3        //2
)

//跳过某些值
const (
	n1 = iota //0
	n2        //1
	_
	n4        //3
	n3        //4
)

// iota 以行计算，同一行之内 仅为0
const (
	b1 = iota //0
	b2 = 100  //100
	b3 = iota //2
	b4        //3
)
const n5 = iota //0



// iota 以行计算，同一行之内 仅为0
const (
	d1,d2 = iota + 1 ,iota + 2 // d1=1 ,d2=2
	d3,d4 = iota + 1 ,iota + 2 // d3= 2,d4=3

)
func main() {

	//fmt.Println("n1:",n1)
	//fmt.Println("n2:",n2)
	//fmt.Println("n3:",n3)
	//
	//fmt.Println("a1:",a1)
	//fmt.Println("a2:",a2)
	//fmt.Println("a3:",a3)
}
