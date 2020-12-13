package main

import (
	"errors"
	"fmt"
)


func divNumber(a,b int) int {
	defer func() {
		if err:= recover();err!=nil{
			fmt.Println(err)
		}
	}()
	return a/b
}

func f1()(r int){
	defer func() {
		r++
	}()
	r = 0
	return
}

//func double(a int)(r2 int) {
//	return a + a
//}

func triple(a int)(r int) {
	defer func() {
		r += a
	}()
	return double(a)
}

func main()  {
	err := errors.New("这是文件的开始")
	fmt.Println(err.Error())
	err1 := fmt.Errorf("%s", "这是一个错误的打印")
	fmt.Println(err1)
	divNumber(100, 0)
	fmt.Println("文件结束")

	f1r := f1()
	fmt.Println(f1r)

	c := triple(3)
	fmt.Println(c)
	fmt.Println(r)
}
