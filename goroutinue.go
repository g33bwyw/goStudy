package main

import (
	"fmt"
	"runtime"
)

/**
本demo主要是学习如何使用go进行并发
当go主程结束，go程也相应的结束

同时了解runtime的相关使用
runtime.Gosched让出当前使用的cpu执行权限，让下一个程序执行。等下次执行的时候再从上一次执行的过程中恢复过来
runtime.Goexit()当前go程终止进行并且退出，不会继续往下走
*/

func testGoExit() {
	defer fmt.Println("1111111111")
	fmt.Println("22222222222")
	runtime.Goexit()
	fmt.Println("333333333") //此处不会被打印

}
func main() {
	//runtime.GOMAXPROCS(1)		//设置使用的cpu核数
	//fmt.Println("----aaaaaaaaaaaaaaa----------")
	//go func() {
	//	for  {
	//		fmt.Println("11111111111111111")
	//		time.Sleep(time.Microsecond)
	//	}
	//
	//}()
	//
	//for  {
	//	fmt.Println("bbbbbbbbbb")
	//	time.Sleep(time.Microsecond)
	//}
	//runtime.GOMAXPROCS(1)

	/**
	当使用gosched go2永远是在最后
	*/
	//go func() {
	//	for i:=0;i<=5;i++{
	//		fmt.Println("我是go1：", i)
	//		i++
	//		//time.Sleep(time.Second)
	//	}
	//}()
	//
	//
	//for  u:=0;u<=2;u++{
	//	runtime.Gosched()
	//	fmt.Println("我是go2:",u)
	//}

	fmt.Println("aaaaaaaa")
	go testGoExit()
	fmt.Println("bbbbbbbbb")
	for {

	}
}
