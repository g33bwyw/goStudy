package main

import "fmt"

func main() {
	//var num1 int
	//var num2 int = 40
	//var avg float64
	//num1 = 35
	//avg = float64((num1 + num2)/2)
	//fmt.Printf("%f", avg)

	var num int
	fmt.Println("请输入一个数字")
	fmt.Scan(&num)
	switch {
	case num>=10 && num <20:
	fmt.Println("low")
	case num>=20:
		fmt.Println("middle")
		fallthrough
	case num>=30:
		fmt.Println("middle")
	default:
		fmt.Println("???????????")

	}

}