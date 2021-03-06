package main

import "fmt"

func remove(s []int, i int) []int {
	slen := len(s)
	copy(s[i:], s[i+1:slen])
	return s[:slen-1]
}

func main() {
	//var s []int = []int{1,2,3,4,5,6,7,8,9,0}
	//fmt.Println(s[:])
	//s1 := s[2:5]
	//s1[2] = 888
	//s2 := s1[2:7]
	//s2[2] = 999
	//fmt.Println(s1, cap(s1),cap(s2))

	slice1 := []int{2, 3, 4}
	slice2 := []int{4, 5, 6, 7}
	slice3 := copy(slice1, slice2)

	fmt.Println(slice3, slice2, slice1)

	s := make([]int, 3)
	a := append(s, 1, 3, 5)
	fmt.Println(a)

	fmt.Printf("%p", s)

	fmt.Println("1111111111111")
	slice4 := []int{5, 6, 7, 8, 9}
	//a3 := remove(slice4 , 1)
	//fmt.Println(a3)

	slice5 := slice4[:2]
	fmt.Println(slice5)
	slice6 := slice4[2:5]
	fmt.Println(slice6)
	copy(slice5, slice6)
	fmt.Println(slice6)

	//var slice8 = make([]int, 0)
	//slice8 = append(slice8,123)
	//fmt.Println(slice8)

}
