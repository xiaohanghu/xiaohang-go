package encapsulation

import (
	"fmt"
)

type Cat struct {
	Animal // 匿名结构体, 相当于继承Animal
	animal Animal // 实名结构体，相当于组合
	// *Animal 也可以用指针形式
	// 可以继承多个父类，如果多个父类里有同名字段，则子类在操作时必须指定访问的是哪个父类的字段 
	Name string // 相同的属性会在子类和父类里各自占用空间
	int // 基本类型也可以匿名, cat.int即可访问。但是因为首字母是小写，不可跨包访问
}

// 父类的方法可以被覆盖 
func (obj *Cat) SetWeight(weight int) {
	fmt.Printf("Cat.SetWeight(%v) \n", weight)
	obj.weight = weight
}

// 类方法不可以重载
// func (obj *Cat) SetWeight(w1 int, w2 int) {
// 	fmt.Printf("Cat.SetWeight(%v, %v) \n", w1, w2)
// 	obj.weight = (w1 + w2)
// }