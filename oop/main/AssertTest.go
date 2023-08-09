package main

import (
	"fmt"
	// domain "oop/domain"
	"reflect"
)

type Device1 interface {
}

type Usb1 interface {
	Device1
	// 接口中不能声明常量
	start()
	stop()
}

type Phone1 struct {
}

type Computer1 struct {
}

func (obj *Phone1) start() { // 注意：*Phone意味着是这个指针类型实现了接口，而Phone类并没有实现接口
	fmt.Printf("Phone.start() \n")
}

func (obj *Phone1) stop() {
	fmt.Printf("Phone.stop() \n")
}

func (obj *Computer1) start() { // 注意：*Phone意味着是这个指针类型实现了接口，而Phone类并没有实现接口
	fmt.Printf("Computer.start() \n")
}

func (obj *Computer1) stop() {
	fmt.Printf("Computer.stop() \n")
}

func main() {

	// var phone = &Phone{}
	var usbPhone Usb1 = &Phone1{}
	var usbComputer Usb1 = &Computer1{}

	// 断言 Assert 可以用于变量的接口类型的转换
	phone, isType := usbPhone.(*Phone1) // .({Type}) 相当于一次方法调用, 只有接口变量可以这么调用
	fmt.Printf("phone:%v, isType: %v \n", phone, isType)

	var phone1 *Phone1 = usbPhone.(*Phone1)            // 这种方式如果转化失败会报错 panic: interface conversion: main.Usb is *main.Computer, not *main.Phone
	var computer *Computer1 = usbComputer.(*Computer1) // 这种方式如果转化失败会报错

	fmt.Printf("phone:%v \n", phone1)
	fmt.Printf("computer:%v \n", computer)

	// typeOfObj := usbPhone.(type) // 只能在switch里用.(type). invalid syntax tree: use of .(type) outside type switch
	typeOfObj := reflect.TypeOf(usbPhone)
	fmt.Printf("type of typeOfObj:%T, typeOfObj:%v, Kind():%v \n", typeOfObj, typeOfObj, typeOfObj.Kind())
	// Type 是类型。Kind 是类别

	switch usbPhone.(type) {
	case *Phone1:
		fmt.Printf("This is a *Phone. \n")
	case *Computer1:
		fmt.Printf("This is a *Computer. \n")
	}

	_, isType = usbPhone.(Device1) // true
	fmt.Printf("Usb is Device: %v \n", isType)
	_, isType = usbPhone.(*Phone1) // true
	fmt.Printf("Usb is *Phone: %v \n", isType)
}
