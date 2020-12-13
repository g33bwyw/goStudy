package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s1 := " hello word"
	//1.字符串是否包含
	fmt.Println(strings.Contains(s1, "wo"))
	//2.去除首尾字符串
	fmt.Println(strings.Trim(s1," "))
	//3.字符串分割
	fmt.Println(strings.Split(s1, "e"))
	//4.字符串重复
	fmt.Println(strings.Repeat("go", 4))
	//5.字符串连接
	slice1 := []string{"安徽", "合肥", "巢湖"}
	fmt.Println(strings.Join(slice1, "-"))
	//6.查找字符串所在的位置
	fmt.Println(strings.Index(s1,"w"))
	//7.去除收尾空格，并且以空格将字符串修改成切片
	fmt.Println(strings.Fields(s1))
	//8.字符串进行替换
	fmt.Println(strings.Replace(s1,"l", "c", 1))

	//将其他类型的转化为字符串strconv.Formate系列函数
	s2 := true
	fmt.Printf("%T:%s\n", strconv.FormatBool(s2), strconv.FormatBool(s2))

	s3 := 1588
	fmt.Printf("%T:%\n", strconv.Itoa(s3), strconv.Itoa(s3))

	s4 := "true"
	s5,_ := strconv.ParseBool(s4)

	fmt.Printf("%T:%t\n", s5,s5)


	s6 := "1688"
	s7,_ := strconv.Atoi(s6)
	fmt.Printf("%T:%d\n", s7,s7)

	//append追加，新建一个字节的切片(讲切片转化为字符串)
	content := make([]byte, 0, 1024)
	content = strconv.AppendBool(content, true)
	content = strconv.AppendInt(content,156, 10)
	content = strconv.AppendQuote(content, "hello")
	fmt.Println(string(content))

	//content2 := []byte{'a', 'b'}
	content2 := make([]byte, 2)
	content2[0] = 'a'
	content2[1] = 'b'
	fmt.Println(content2)
	//content2 = append(content2, 'a')
	fmt.Println(string(content2))



}
