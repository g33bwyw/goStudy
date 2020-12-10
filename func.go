package main

import "fmt"

var a int
type Add func(a,b int) int

func getAddRes(a,b int) int {
	return a + b
}
func test() {
	a = 5;
	a++
	fmt.Println("test1", a)
}

func getArgs(args ...int) int {
	arr := args[2:]
	fmt.Printf("%T", arr)
	fmt.Println(arr)
	var sum int
	for _,v := range args{
		sum += v
	}
	return sum
}
func main() {
	//局部变量
	a := 9
	test()
	fmt.Println("main:",a)

	var c  Add
	c = getAddRes
	fmt.Println(c(7,2));

	d := getArgs(8,3,5,11)
	fmt.Println(d)

	//匿名函数
	var lastRes int = 2
	e := func(a,b int) int{
		lastRes := 0
		lastRes ++
		fmt.Println("func lastRes", lastRes)
		return a + b
	}(1 ,2)
	fmt.Println(e)
	fmt.Println(lastRes)

}
