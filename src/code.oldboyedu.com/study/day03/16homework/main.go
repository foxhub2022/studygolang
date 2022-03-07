package main

import (
	"fmt"
	"strings"
)

/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/
var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func main() {
	//left := dispatchCoin()
	//fmt.Println("剩下：", left)


	left := dispatchCoin()
	fmt.Println("剩下：", left)
}

//a. 名字中每包含1个'e'或'E'分1枚金币
//b. 名字中每包含1个'i'或'I'分2枚金币
//c. 名字中每包含1个'o'或'O'分3枚金币
//d: 名字中每包含1个'u'或'U'分4枚金币
// 此方法用string index方法来判断，如果存有响应数据 则加入对应参数
func checkName (name  string )  int {
	//返回用于硬币
	getCoin := 0

	upperName := strings.ToUpper(name)

	hasE := strings.Index(upperName,"E")
	haseI := strings.Index(upperName,"I")
	haseO := strings.Index(upperName,"O")
	hasU := strings.Index(upperName,"U")

   if 	hasE > 0 {
	   getCoin += 1
   } else if haseI > 0 {
	   getCoin += 2
   }  else if haseO > 0 {
	   getCoin += 3
   } else if hasU  > 0  {
	   getCoin += 4
   }
   //fmt.Println("得到金币",getCoin)
   return  getCoin
}


func dispatchCoin()	int  {
	// 1.依次拿到每个人的名字
	// 2.拿到一个人名根据分金币的规则区分金币
	// 2.1 每个人区分的金币数量应该保存到 distribution中
	// 2.2 还要记录下剩余的金币数量
	// 3 整个2步执行就能得到最终每个人的金币数量和剩余金币数
	leftConin := 0
	for _, name := range users {
		//fmt.Println(name)
		totalCoin :=  checkName(name)
		fmt.Printf("name: %s Coin：%d  \t \n",name,totalCoin)
		distribution[name] = totalCoin //将名字和获得硬币保存到map
		coins -= totalCoin  //剩余金币
	}
	leftConin = coins
	return leftConin
}