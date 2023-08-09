package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	fmt.Printf("%v, type: %T\n", now, now)
	fmt.Printf("%v-%v-%v %v:%v:%v\n", now.Year(), int(now.Month()), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println(now.Format("Mon Jan 2 15:04:05 -0700 MST 2006"))
	fmt.Println(now.Format("2006-01-02 15:04:05"))
}
