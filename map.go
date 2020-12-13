package main

import "fmt"

func main() {
	var a map[string]string = map[string]string{
		"a" : "1",
		"b" : "2",
	}
	c := a
	c["a"] = "678"
	fmt.Println(a)
	fmt.Println(c)

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

}
