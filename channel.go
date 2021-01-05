package main

import (
	"fmt"
	"time"
)

func printer(s string) {
	for _,ch := range s {
		fmt.Printf("%c", ch)
		time.Sleep(time.Second)
	}
}
var channel = make(chan int)
func person1() {
	str := "hello"
	printer(str)
	channel<-10
}

func person2() {
	<-channel
	str := "world"
	printer(str)
}
func main() {

	go person1()
	go person2()
	for {
		;
	}

}
