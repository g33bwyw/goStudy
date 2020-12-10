package main

import "fmt"

func main() {
	const (
		a = 10
		b = iota

		d = 20
		c = iota
	)
	fmt.Println(a,b,c,d)

	const (
		e,f,j = iota,iota,iota
		m = iota
	)
	fmt.Println(e,f,j,m)
}