package _3bibao

import "fmt"

func calc(index string,a,b int) int {
	ret := a + b
	fmt.Println(index,a,b,ret)
	return ret
}

func main() {
	a := 1
	b := 2
	defer calc("1",a,calc("10",a,b) )
	a = 0
	defer calc("2",a,calc("20",a,b) )
	b = 1
}


// 1. defer calc ("1",a,calc("10",a,b))
// 2. calc("10",a,b) // "10" 1 2 3
// 3. defer calc("1",1,3)
// 4. a = 0
// 5. defer calc("2",a calc("20",a,b))
// 6. calc("20",a,b) // "20" 0 2 2
// 7. defer calc("2",0,2)
// 8. b=1
// 9. calc("2",a,2) // "2" 0 2 2
// 10.calc("1",a,3) // "1" 1 3 4


// "10" 1 2 3
// "20" 0 2 2
// "2" 0 2 2
// "1" 1 3 4