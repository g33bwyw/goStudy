package main

import "fmt"

func main() {
	for i:=1;i<= 9;i++ {
		for j:=1;j<= 9;j++ {
			if (i < j) {
				fmt.Printf("\t")
			} else {
				fmt.Printf("%d*%d=%d\t",i,j,i*9)
			}
		}
		fmt.Println()
	}


}