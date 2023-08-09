package test

import (
	"fmt"
	"testing"
	"time"
)

func Test_channel(t *testing.T) {
	//var intChan chan  int
	intChan := make(chan int, 3)
	// 放入的值超过channel大小则会报错：fatal error: all goroutines are asleep - deadlock!

	//intChan <- 13

	go func() {
		for {
			num := <-intChan
			fmt.Printf("r1 num: %v \n", num)
		}
	}()

	go func() {
		for {
			num := <-intChan
			fmt.Printf("r2 num: %v \n", num)
		}
	}()

	for num := 4; num <= 100; num++ {
		intChan <- num
	}

	time.Sleep(time.Second * 1)
}
