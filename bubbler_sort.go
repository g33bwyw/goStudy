package main

import "fmt"

//冒泡排序【4,5,6,3,2,1】
func bubble_sort(arr [6]int) [6]int{
	arrLen := len(arr)
	for i:=0;i<arrLen;i++ {
		for j:=0;j<arrLen-i-1;j++{
			if arr[j] > arr[j+1] {
				arr[j],arr[j+1] = arr[j+1],arr[j]
			}
		}
	}
	return arr
}

func main() {
	var arr [6]int = [6]int{4,5,6,3,2,1}
	arr2 := bubble_sort(arr)
	fmt.Println(arr2)
}



