package main

import "fmt"


//冒泡排序
func sortByMain(arr [10]int) [10]int {

	length := len(arr)
	for i:=0;i<length;i++ {
		for j:=0;j<length-i-1;j++ {
			if arr[i] > arr[j] {
				arr[j],arr[i] = arr[i],arr[j]
			}

		}
	}

	return arr
}
func main() {
	//var a [10]int = [10]int{7,8,11,9,68,200,45,2,35,85}
	a := [10]int{7,8,11,9,68,200,45,2,35,85}
	var maxNum int
	var minNum int
	var avgNum int
	var sumNum int
	fmt.Println(a)
	for i,v := range a {
		if i == 0 {
			minNum = v
		}
		sumNum += v
		if maxNum < v {
			maxNum = v
		}
		if minNum > v {
			minNum = v
		}
	}
	avgNum = sumNum/10

	fmt.Println(maxNum,minNum,avgNum,sumNum)

	sortA := sortByMain(a)
	fmt.Println(sortA)
}