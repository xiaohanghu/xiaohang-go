

package main

import (
	"fmt"
	// domain "oop/domain"
)

type Device interface {
	start(n int) (int)
	stop()
}

type Usb interface {
	// 接口中不能声明常量
	Device // 接口 可以继承多个接口 
	SetState(s int)
}

type Phone struct { 
	// 不需要显示实现接口，只需要实现接口一样的方法
	// 必须实现接口的所有方法，否则无法当作接口的类型传入方法
	State int
}

func (obj *Phone) start(n int) int { // 注意：*Phone意味着是这个指针类型实现了接口，而Phone类并没有实现接口
	fmt.Printf("Phone.start(%v) \n", n)
	return n + 1 
}

func (obj *Phone) stop() {
	fmt.Printf("Phone.stop() \n")
}

func (obj *Phone) SetState(s int) {
	obj.State = s
}

func (obj *Phone) String() string {
	return fmt.Sprintf("{State: %v}", obj.State)
}

func (obj *Phone) Call() { // 注意：*Phone意味着是这个指针类型实现了接口，而Phone类并没有实现接口
	fmt.Printf(".. Phone.Call() \n")
}

func work(usb Usb){ // 这里的Usb 可以具体是值类型，也可以是指针
	fmt.Printf("usb addr2: %p \n", usb)
	fmt.Printf("usb addr2-1: %p \n", &usb)
	n := usb.start(3)
	fmt.Printf("work() return %v , type: %T \n", n, usb)
}

func main(){

	// var phone = &Phone{}
	var usb Usb = &Phone{}
	fmt.Printf("usb addr1: %p \n", usb)
	fmt.Printf("usb addr1-1: %p \n", &usb)
	usb.SetState(7)
	work(usb)
	fmt.Printf("usb: %v \n", usb)

	fmt.Println()
	var usbs = []Usb{&Phone{State: 3}, &Phone{}, &Phone{}}
	fmt.Println("usbs:", usbs)
	fmt.Println()


	// var empty interface{} = usb
	var empty interface{} = 3.4 // interface{} 代表空接口, 任何类型都可以赋 给空接口
	fmt.Printf("empty: %v \n", empty)

} 