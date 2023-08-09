package main

import (
	"fmt"
)

// go env -w GO111MODULE=on
// go env -w GOPATH=/Users/xiaohanghu/Documents/Repositories/xiaohang-go
func main() {

	v1 := "name"
	n1 := 12
	switch v1{
		case "name":
			fmt.Println(`case "name"`)
			// 默认带了break
			fallthrough
		case "bbb", "cc":
			fmt.Println(`case "bbb", "cc"`)
		default:
			// fmt.Println(`default`)
	}

	n1 := 12
	switch {
		case n1>10:
			fmt.Println(`case n1>10`)
	}
}
