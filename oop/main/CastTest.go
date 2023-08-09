

package main

import (
	"fmt"
)

type Teacher struct{
	Age int
}

type Person struct{
	Age int
}

func main(){
	var t Teacher = Teacher{3}
	var p Person= Person{11}
	t = Teacher(p) // 只有字段名字完全一样才可以做强转赋值
	fmt.Println("t:", t)
	fmt.Printf("&t: %p \n", &t)
	fmt.Printf("&p: %p \n", &p)
} 