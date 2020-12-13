package main

import "fmt"

func main()  {
	var s map[string]string
	s = make(map[string]string)
	s["id"] = "100"
	s["name"] = "wyw"
	fmt.Println(s)

	delete(s, "id")
	fmt.Println(s)

}
