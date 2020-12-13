package main

import "fmt"

func insert_sort(arr [6]int) [6]int{
	arrLen := len(arr)
	for i:=0;i<arrLen;i++ {
		
	}
	return arr
}

func main() {
	var arr [6]int = [6]int{4,5,6,3,2,1}
	arr2 := insert_sort(arr)
	fmt.Println(arr2)
}