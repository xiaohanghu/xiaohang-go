package main

import (
	"fmt"
	// "unsafe"
	// "unicode/utf8"
	// "strings"
)


func main() {
	// map 是内置类型, key无序
	// var kv map[int]string = make(map[int]string)
	var kv map[int]string = make(map[int]string, 1) // 这里是初始大小，map可以自动扩容
	kv[1] = "safasf"
	kv[3] = "dd"

	fmt.Printf("&kv: %v \n", &kv)

	kv[14] = "rc"
	fmt.Printf("kv: %v \n", kv)

	kv = map[int]string{
		1: "张三",
		7: "李四",
		41: "王五",
	}
	delete(kv, 7)
	fmt.Printf("kv: %v \n", kv)

	var v, flag = kv[41]
	fmt.Printf("v: %v, flag:%v \n", v, flag)

	fmt.Printf("len(kv):%v \n", len(kv))

	for k,v := range kv {
		fmt.Printf("%v:%v \n", k, v)
	}

	map1 := map[int]map[int]string{
		1:{
			3:"hah",
		},
		5:{
			9:"test",
			30:"odsds",
		},
	}
	fmt.Println("map1:\n", map1)

}
