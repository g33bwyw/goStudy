package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getRand(num int) int {

	return rand.Intn(num)
}

func sumRand(total int, num int) int {
	total = total + num
	return total
}

func searchNum(arr [5]int) {
	for i, v := range arr {
		for m, n := range arr {
			if m > i && arr[i]+arr[m] == 8 {
				fmt.Println("查找出来指定元素：", i, m, "值：", v, n)
			}
		}
	}
}

func main() {
	var Arr [5]int
	var total int
	fmt.Println(Arr)
	rand.Seed(time.Now().Unix())
	for i, _ := range Arr {
		num := rand.Intn(10)
		Arr[i] = num
		total = sumRand(total, num)
	}
	fmt.Println(Arr)
	fmt.Println(total)

	searchNum(Arr)
}
