package main

import (
	"fmt"
)

// go env -w GO111MODULE=on
// go env -w GOPATH=/Users/xiaohanghu/Documents/Repositories/xiaohang-go
func main()  {

	var age int
	var name string

	// fmt.Println("请录入学生的年龄：")
	// fmt.Scanln(&age) //以换行作为结束

	// fmt.Println("请录入学生的名字：")
	// fmt.Scanln(&name)

	fmt.Println("请录入学生的年龄 名字：")
	fmt.Scanf("%d %s", &age, &name)
	fmt.Println("age：", age, ", name:", name)
}