package main

import "fmt"

func removeSlice(slice []int, i int) {
	slice1 := slice[:i]
	slice2 := slice[i+1:]
	slice3 := append(slice1, slice2...)
	fmt.Println(slice3)
}

func removeSlice2(slice []int, i int) {
	copy(slice[i:], slice[i+1:])
	fmt.Println(slice[:len(slice)-1])
}

func main() {
	slice := []int{1,2,3,4,5}
	removeSlice2(slice, 2)

}