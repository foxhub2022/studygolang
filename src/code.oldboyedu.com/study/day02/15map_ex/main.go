package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var scoreMap = make(map[string]int,200)

	for i := 0; i < 100;i++ {
		key := fmt.Sprintf("stu%02d",i)
		value := rand.Intn(100)
		scoreMap[key] = value
	}

	var keys = make([]string,0,200)
	for key :=range scoreMap {
		keys = append(keys,key)
	}
	fmt.Println("***********************************************")
	fmt.Println(keys)
	fmt.Println("***********************************************")

	sort.Strings(keys)
	for _,key := range keys {
		fmt.Println(key,scoreMap[key])
	}
}
