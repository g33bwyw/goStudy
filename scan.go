package main

import "fmt"

func main() {
	var age string
	var name string
	fmt.Println("请输入您的姓名")
	fmt.Scanf("%s", &name)
	fmt.Println("请输入您的年龄")
	fmt.Scanf("%s", &age)
	fmt.Println(name, age)

	//var ch byte
	//ch = 'a'
	//fmt.Printf("%c", ch)


}