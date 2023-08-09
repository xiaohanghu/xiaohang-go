package main

import (
	"fmt"
	"unsafe"
	// "unicode/utf8"
)

/*
bool 默认值：false
string
int(4 or 8 bytes)	int8(1 byte)	int16(2 bytes)	int32(4 bytes)	int64(8 bytes)
uint				uint8			uint16			uint32			uint64			uintptr 无符号
byte // alias for uint8 [0,255]
rune // alias for int32
     // represents a Unicode code point
float32(4 bytes)	float64(8 bytes)
complex64 complex128

字符使用UTF-8编码
*/

// 全局变量，不能使用  :=
var v1 = 10
var (
	v2 = 20.3
	v3 = "hi"
)

func main() {
	var age byte
	age = 10
	var age2 int = 20
	// var age2 int = 19 //error
	var num int // num = 0

	// :=”只能在声明“局部变量”的时候使用，而“var”没有这个限制
	sex := "男"

	fmt.Println("age = ", age)
	fmt.Println("age2 = ", age2)
	fmt.Println("num = ", num)
	fmt.Println("sex = ", sex)

	n1, n2, n3 := 4, 314e-2, "101"
	fmt.Printf("n1的类型是%T, 占 %d bits", n1, unsafe.Sizeof(n1))
	fmt.Println()
	fmt.Printf("n1=%d, n2=%6.2f, n3=%s\n", n1, n2, n3)

	fmt.Printf("v1=%d, v2=%6.2f, v3=%s\n", v1, v2, v3)

}
