package main

import (
	"fmt"
)

func main() {

	var p *int = new(int) // 开辟类型对应的内存并返回其地址
	*p = 21
	fmt.Println(*p)
	
}
