package main

import (
	"errors"
	"fmt"
)

func returnValues() int {
	var result int
	defer func() {
		result++
		fmt.Println("defer return")
	}()
	return result // 返回0
}

func namedReturnValues() (result int) { // 这种方式相当于方法自己开辟了 result 变量的空间
	// var result int
	defer func() {
		result++
		fmt.Println("defer named return")
	}()
	// 返回1
	return result // 返回1
}

func testRecover() {
	defer func() {
		fmt.Println("~~~defer")
		err := recover()
		if err != nil {
			fmt.Println("~~~err:", err)
		}
	}()

	n1 := 10
	n2 := 0
	// result := 10 / 0 // 这种会编译错误
	result := n1 / n2
	fmt.Println(result)
}

func test1(n1 int, n2 int) (int, error) {
	if n2 == 0 {
		return -1, errors.New("Error: n2 == 0")
	}
	return n1 / n2, nil
}

func main() {
	fmt.Println("returnValues:", returnValues())
	fmt.Println("namedReturnValues:", namedReturnValues())

	testRecover()

	r, err := test1(5, 0)
	fmt.Println("r:", r, ", err:", err)

}
