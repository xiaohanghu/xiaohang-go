package main

import (
	"fmt"
)

// Closure 
func getSum() func (int) (int) { //getSum的返回类型为一个函数
	var sum int = 0
	return func (n int) (int) {
		sum += n
		return sum
	}
}

// Closure 
var getId func() int = func () func() int {
	var id int = 1000
	return func () int {
		id++
		return id
	}
}()

func main() {

	n := func (a int, b int)(int){
		return a + b
	}(5, 7)

	fmt.Println("n:", n)

	sub := func (a int, b int)(int){
		return a - b
	} 

	fmt.Println("sub:", sub(10, 7))

	// Closure
	a := 10
	n1 := func (b int)(int){
		return a + b
	}(7)
	fmt.Println("n1:", n1)


	f := getSum()
	f(1)
	fmt.Println("f:", f(2))

	fmt.Println("getId:", getId())
	fmt.Println("getId:", getId())

}
