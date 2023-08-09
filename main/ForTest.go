package main

import (
	"fmt"
)

func main() {

	// 没有while do..while, 只有for
	sum := 0
	for i := 1; i <= 5; i++ {
		sum += i
	}
	fmt.Println("sum:", sum)

	var i int = 0
	for i < 10 {
		fmt.Println("0-i:", i)
		i++
	}
	fmt.Println()
	i = 0
	for {
		fmt.Println("1-i:", i)
		i++
		if i >= 10 {
			break
		}
	}

	fmt.Println()

	lable2: // 直接跳出外层循环
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			fmt.Println("i:", i, ",j:", j)
			if i==2 && j==2{
				// continue
				break lable2
			}
		}
	}
}
