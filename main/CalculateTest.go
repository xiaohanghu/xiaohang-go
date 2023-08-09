package main

import (
	"fmt"
	// "unsafe"
	// "unicode/utf8"
)

func main(){
	var n1 int = +10 //正10
	fmt.Println(n1)


	fmt.Println(10/3)
	fmt.Println(10./3)

	//  a%b=a-(a/b)*b

	var n2 int = 10
	n2++ //++只能在变量后面
	// n3 := n2++ 非法
	fmt.Println(n2)

	a, b := 3, 4
	a, b = b, a
	fmt.Println("a:", a, ",b:", b)


	
}