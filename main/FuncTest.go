package main

import (
	"fmt"
)

func sumAndSub(num1 int, num2 int) (int, int) {
	return num1 + num2, num1 - num2
}

func sum(args... int) (int) {
	var sum int = 0
	for i:=0; i< len(args); i++ {
		sum += args[i]
	}
	return sum
}

func changeValue(v *int){
	*v = 10
}

func testFunc(v1 int, v2 int, callback func(int, int)){
	v1++
	v2++
	callback(v1, v2)
}

func sumAndSub01(num1 int, num2 int) (sum int, sub int) {
	sum = num1 + num2
	sub = num1 - num2
	return
}

func main() {
	s,d := sumAndSub01(3, 5)
	fmt.Println("sumAndSub01:", s, d)

	fmt.Println("sum:", sum(1,2,3,4))

	var v int = 4
	changeValue(&v)
	fmt.Println("v:", v)

	f := changeValue 
	fmt.Printf("f的类型是: %T \n", f)

	testFunc(1, 2, func(v1 int, v2 int){
		fmt.Printf("v1:%v, v2:%v \n", v1, v2)
	})

	type myInt int //自定义类型 
}
