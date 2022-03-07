package main
import (
	"fmt"
)

func main() {
	fmt.Println("start")
	for i:=1;i<10;i++{
		//fmt.Printf("%d \t",i)
		for j := 1;j <= i ;j++ {
			fmt.Printf("%d * %d = %d \t",i,j,i*j)
		}
		fmt.Println()
	}

	a := 1
    a--
    a = 2



	fmt.Println("stop")

	//XX:
	//	fmt.Println("over")
}

