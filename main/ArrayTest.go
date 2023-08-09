package main

import (
	"fmt"
	"reflect"
)

func testSetValue(arr [5]int16) { //数组是一定要指定长度的，否则就是slice
	arr[0] = 100
}

func main() {
	// var arr [3]int16  //默认值0
	// var arr = []int16{4, 2, 6}
	var arr = [5]int16{2: 4, 1: 2, 3: 6, 4: 43}
	var slice = []int16{2: 4, 1: 2, 3: 6, 4: 43}
	fmt.Println("arr .Kind(): ", reflect.TypeOf(arr).Kind())
	fmt.Println("slice.Kind(): ", reflect.TypeOf(slice).Kind())
	// arr[1] = 7
	// 长度是类型的一部分

	// for i:=0; i< len(arr); i++ {
	// 	fmt.Println("请输入:")
	// 	fmt.Scanln(&arr[i])
	// }

	testSetValue(arr) // 数组属于值类型而非引用类型，在作为方法参数时， 会拷贝整个数组

	fmt.Println("arr:", arr)
	fmt.Println("cap(arr):", cap(arr))
	for i, v := range arr {
		fmt.Println(i, ":", v)
	}

	fmt.Println("-----------")
	var arr2 = [][]int{{1, 2}, {5, 6, 7}}
	fmt.Println("arr2[1][1]:", arr2[1][1])
	for i, v := range arr2 {
		fmt.Println(i, ":", v)
	}
}
