package main

import (
	"fmt"
	"strings"
)

func removeMap(a map[string]string, i string) {
	delete(a, i)
}

func statWord(s string) {
	slice := strings.Fields(s)
	fmt.Println(slice)
	var statMap map[string]int = map[string]int{}
	for _, v := range slice {
		statMap[v] = statMap[v] + 1
	}
	fmt.Println(statMap)
}
func main() {
	var a map[string]string = map[string]string{
		"a": "1",
		"b": "2",
		"c": "3",
	}
	removeMap(a, "b")
	fmt.Println(a)
	//c := a
	//c["a"] = "678"
	//fmt.Println(a)
	//fmt.Println(c)

	//for k,v := range a {
	//	fmt.Printf("%s=>%s\t", k, v)
	//}
	//
	//b := [3]string{"a", "b", "c"}
	//for k1,v1 := range b {
	//	fmt.Printf("%d=>%s\t", k1, v1)
	//}
	//
	////删除元素
	//delete(a, "b")
	//fmt.Println(a)

	var word string = "I love my work and I love my family too"
	statWord(word)
}
