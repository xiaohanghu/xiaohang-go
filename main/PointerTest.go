package main

import (
	"fmt"
	// "strconv"
)

func main() {
	// https://studygolang.com/pkgdoc

	var v1 int64 = 1045
	var p1 *int64 = &v1 // 获取地址
	// p1 = p1 ++ // 不支持运算
	fmt.Printf("p1:%v \n", p1)
	var p11 **int64 = &p1
	fmt.Printf("p11:%v \n", p11)

	var v11 int64 = *p1 // 获取地址指向的值
	fmt.Printf("v11:%v \n", v11)

	*p1 = 322 // 改变指针指向的值 
	fmt.Printf("v1:%v \n", v1)

}
 