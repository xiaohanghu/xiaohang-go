package main

import (
	"fmt"
	"strconv"
)

func main() {
	// https://studygolang.com/pkgdoc

	var v1 int64
	v1, _ = strconv.ParseInt("54311", 10, 64) // _是个特殊关键字，可以不申明而直接使用

	// var v11 int
	// v11, _ = strconv.Atoi("54311") // Atoi是ParseInt(s, 10, 0)的简写。

	fmt.Println("v1: ", v1)

	var v2 float64
	v2, _ = strconv.ParseFloat("3.1415", 64) // _是个特殊关键字，可以不申明而直接使用
	fmt.Println("v2: ",v2)

}
 