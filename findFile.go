package main

import (
	"fmt"
	"os"
)

/**
检查文件是否是文件
*/
func getFile(filepath string, fileExt string) ([]string, error) {
	var fileCollect []string = []string{}
	file, err := os.OpenFile(filepath, os.O_RDONLY, os.ModeDir)
	if err != nil {
		return fileCollect, fmt.Errorf("%s文件不存在", filepath)
	}
	basePath, _ := os.Getwd()
	fileList, _ := file.Readdir(-1)
	for _, fileDetail := range fileList {
		fmt.Println(basePath)
		fmt.Println(fileDetail.Name())
		filepath = fmt.Sprintf("%s/goStudt/%s", basePath, fileDetail.Name())
		if fileDetail.IsDir() {
			getFile(filepath, fileExt)
		} else {
			fileCollect = append(fileCollect, filepath)
		}
	}

	return fileCollect, nil
}


func main() {
	var str string = "d:/b.txt"
	f, err := os.OpenFile(str, os.O_RDWR, 6)
	fmt.Println(f, err)
	//var source string
	//var fileExt string
	//fmt.Println("请输入想要找的目录:")
	//fmt.Scanf("%s", &source)
	//fmt.Println("请输入想要文件扩展名:")
	//fmt.Scanf("%s", &fileExt)
	////fmt.Println(source, fileExt)
	//fileSlice, err := getFile(source, fileExt)
	//if err != nil {
	//	fmt.Println(err.Error())
	//} else {
	//	fmt.Println(fileSlice)
	//}
}
