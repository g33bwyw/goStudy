package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/**
检查文件是否存在
 */
func checkFileIsExist(filepath string) bool {
	_,existErr := os.Stat(filepath)
	return !os.IsNotExist(existErr)
}

/**
复制文件
 */
func CopyFile(source string, target string) error{
	if !checkFileIsExist(source) {
		return fmt.Errorf("%s文件不存在", source)
	}

	sourceFile, err := os.Open(source)
	if err != nil {
		return fmt.Errorf("%s打开错误", source)

	}
	if checkFileIsExist(target) {
		os.Remove(target)
	}
	targetFile,err2 := os.OpenFile(target,os.O_RDWR | os.O_CREATE, os.ModePerm)
	if err2 != nil {
		return fmt.Errorf("%s文件创建失败", target)
	}
	defer sourceFile.Close()
	defer targetFile.Close()

	r := bufio.NewReader(sourceFile)
	for  {
		content, err3 := r.ReadBytes('\n')
		if err3 != nil {
			if err3 == io.EOF {
				break
			}
			return fmt.Errorf("%s读取文件出现异常", sourceFile)
		}
		targetFile.Write(content)
	}

	return nil
}

func main() {
	args := os.Args
	if args == nil || len(args) != 3 {
		fmt.Println("请输入有效的参数")
		return
	}
	source := args[1]
	target := args[2]
	err := CopyFile(source, target)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(source,"复制成功")
}
