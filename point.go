package main

import "fmt"

func changeArr(arr *[5]int) {
	arr[1] = 100
}

func test() {
	p := new(int)
	*p = 3000
}
func main() {
	////指针定义
	//var i int = 100
	//var p *int
	//p = &i
	//*p = 80
	//fmt.Println(i, *p)
	//
	////指针数组
	//var arr [5]int = [5]int{1,2,3,4,5}
	//changeArr(&arr)
	//fmt.Println(arr)
	//
	////数组指针
	//var arr2 [5]*int
	//var k1 int = 100
	//var k2 int = 200
	//arr2[0] = &k1
	//arr2[1] = &k2
	//fmt.Println(arr2)
	//fmt.Println(arr[1])
	//
	////结构体
	//type Student struct {
	//	id int
	//	name string
	//}
	//stu := &Student{100, "wyw"}
	//fmt.Println(stu)
	//stu2 := stu
	//(*stu2).id = 200
	//fmt.Println(stu)
	//fmt.Println(stu2)
	var p *int
	test()
	fmt.Println(*p)

}
