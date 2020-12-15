package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/**
创建文件,字符串写入
*/
func createFile1(s string) {
	f, err := os.Create(s)
	if err != nil {
		fmt.Println("文件打开错误", err)
		return
	}
	f.WriteString("hello 我这是字符串写入")
	defer f.Close()
}

/**
创建文件，字节写入
*/
func createFile2(s string) {
	f, err := os.Create(s)
	if err != nil {
		fmt.Println("文件打开错误", err)
		return
	}
	str := "hello 我这是字节写入"
	slice := []byte(str)
	f.Write(slice)
	defer f.Close()
}

/**
创建文件，write写入
*/
func createFile4(s string) {
	f, err := os.Create(s)
	if err != nil {
		fmt.Println("文件打开错误", err)
		return
	}
	str := "hello 我这是writeAt写入"
	n, _ := f.Seek(0, os.SEEK_END)
	slice := []byte(str)
	f.WriteAt(slice, n)
	defer f.Close()
}

/**
偏移追加写入
*/

func createFile3(s string) {
	f, err := os.OpenFile(s, os.O_APPEND, 6)
	if err != nil {
		fmt.Println("文件打开错误", err)
		return
	}
	str := "hello 我这是追加写入"
	slice := []byte(str)
	f.Write(slice)
	defer f.Close()
}

/**
文件读取
*/
func readFile1(s string) {
	f, err := os.Open(s)
	if err != nil {
		fmt.Println("文件打开错误", err)
		return
	}
	r := bufio.NewReader(f)
	for {
		buf, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println("err")
			}
		}
		fmt.Println(string(buf))
	}
	defer f.Close()
}

/**
一次文件读取
*/
func readFile2(s string) {
	f, err := os.Open(s)
	if err != nil {
		fmt.Println("文件打开错误", err)
		return
	}
	defer f.Close()
	content := make([]byte, 1024*2)
	n, err2 := f.Read(content)
	if err2 != nil && err2 != io.EOF {
		fmt.Println(err2)
	}
	fmt.Println(string(content[:n]))

}

/**
创建文件已字节写入
*/

func main() {
	//s1 := "d:/a.txt"
	//createFile1(s1)
	//
	//s2 := "d:/b.txt"
	//createFile2(s2)

	//s1 := "d:/a.txt"
	//createFile3(s1)

	//s4 := "d:/c.txt"
	//createFile4(s4)
	s1 := "d:/a.txt"
	//readFile1(s1)
	readFile2(s1)

}
