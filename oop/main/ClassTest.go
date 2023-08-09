package main

import (
	"fmt"
)

// go中struct就是class
type Teacher1 struct {
	name   string // 私有
	Age    int    // 不支持指定默认值
	School string
}

// 该方法绑定了类型Teacher1， 所有的自定义类型都可以用这种方式绑定方法
func (obj *Teacher1) setAge(age int) {
	fmt.Printf("setAge(%v) \n", age)
	// (*obj).Age = age
	obj.Age = age
}

// 非pointer的方式是值传递，无法修改对象的属性
func (obj Teacher1) setSchool(school string) {
	fmt.Printf("setSchool(%v) \n", school)
	obj.School = school
}

// go版toString( )
func (obj *Teacher1) String() string {
	return fmt.Sprintf("{name: %v, Age: %v, School: %v}", obj.name, obj.Age, obj.School)
}

// 这种值拷贝的方式开销较大
// func (obj Teacher1) String() string{
// 	return fmt.Sprintf("{name: %v, Age: %v, School: %v}", obj.name, obj.Age, obj.School)
// }

func setName(t Teacher1, name string) {
	t.name = name
}

func main() {
	var t1 Teacher1             // 这里就已经初始化了, 这里t1是值类型
	fmt.Println("t1 init:", t1) // 字符串的初始值是空
	t1.name = "Dave"
	t1.Age = 30
	t1.School = "UNSW"
	t1.setSchool("###")
	setName(t1, "Li")
	fmt.Println("t1:", &t1)

	var t2 = Teacher1{"Zhao", 33, "HUAKE"} //t2是值类型
	setName(t2, "Li")
	fmt.Println("t2:", t2) // 字符串的初始值是空

	var t3 *Teacher1 = new(Teacher1)
	// var t3 *Teacher1 = &Teacher1{"Zhao", 33, "HUAKE"}
	(*t3).name = "Hang"
	(*t3).Age = 22
	t3.School = "BD" // 指针和值变量都可以直接访问属性
	t3.setAge(101)   // 等于 (*t3).setAge(101), golang类型方法的这2种调用方式是等价的，会根据方法签名上是值还是指针自动转化
	fmt.Println("t3:", t3)

	var t4 *Teacher1 = &Teacher1{ // 这种初始化方式很灵活
		name: "TTT4",
	}
	fmt.Println("t4:", t4)
}
