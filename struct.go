package main

import "fmt"

type Student struct {
	id int
	name string
}

func main() {
	var s Student
	s.id = 1
	s.name = "wangyw"

	s1 := s
	s1.id = 2
	fmt.Println(s)
	fmt.Println(s1)
}




