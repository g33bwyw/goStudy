package main

import "fmt"
type P interface {
	say()
}

type Human interface {
	P
	sing(song string)
}
type Person struct {
	name string
	age int
}

func (people Person)PrintDetail()  {
	fmt.Println("这是people打印：", people)
}

type Student struct {
	Person
	job string
	fav string
}

func (stu *Student) say() {
	fmt.Println("这是interface打印", stu)
}
func (stu *Student) sing(song string) {
	fmt.Println("这是interface继承打印", song)
}

func (stu *Student)EditStudent(name string, age int) {
	stu.age = age
	stu.name = name
}

func (stu Student)PrintDetail()  {
	fmt.Println("这是student打印：",stu)
}

func sumNum(a int, b int) int {
	return a + b
}

func main()  {
	stu := Student{Person{"wyw", 18}, "学习", "basketball"}
	fmt.Println(stu)
	(&stu).EditStudent("jack", 20)
	stu.Person.PrintDetail()
	stu.PrintDetail()
	fmt.Println("-----------------")
	a := (Student).PrintDetail
	a(stu)

	c := sumNum
	d := c(1, 3)
	fmt.Println(d)

	//接口
	var i Human
	i = &stu
	i.say()
	i.sing("php是世界上最好的语言")

	//空接口
	//var emptyInferface []interface{}
	emptyInterface := make([]interface{}, 3)
	emptyInterface = append(emptyInterface, 1)
	emptyInterface = append(emptyInterface, "hello")
	//类型断言
	for _, v := range emptyInterface {
		if _,ok := v.(int);ok== true {
			fmt.Println(v)
		}
	}
	fmt.Println(emptyInterface)
}
