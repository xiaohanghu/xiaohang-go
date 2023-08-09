package main

import (
	"fmt"
	"strconv"
)

func main() {
	var v1 int = 54311
	// Sprintf根据format参数生成格式化的字符串并返回该字符串。
	s1 := fmt.Sprintf("%d", v1)
	fmt.Println(s1)

	// https://studygolang.com/pkgdoc
	// var s11 string = strconv.FormatInt(int64(v1), 16) // 16是base
	var s11 string = strconv.Itoa(v1) // Itoa是FormatInt(i, 10) 的简写。
	fmt.Println(s11)

	var v2 float64 = 4.29
	var s2 string = strconv.FormatFloat(v2,'f', 9, 64) // f表示10进制格式，9表示小数点后保留9位，64表示这个数是float64
	fmt.Println(s2)

	

}
 