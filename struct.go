package main

import "fmt"

type Student struct {
	id int
	name string
}
func main()  {
	type Student struct {
		id int
		name string
	}
	//第一种定义方式
	var s Student = Student{
		100, "wyw",
	}
	fmt.Println(s)
	//第二种定义方式
	var s2  Student
	s2.id = 101
	s2.name = "wyw"
	fmt.Println(s2)

}





