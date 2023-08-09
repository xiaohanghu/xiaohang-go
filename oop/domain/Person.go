package encapsulation

import (
	"fmt"
)

type person struct{
	name string
	age int
}

func NewPerson(name string) *person{
	obj := &person{
		name: "null", // 默认值
		age: -1, // 默认值
	}

	obj.name = name

	return obj
}

func (obj *person) SetName(name string) {
	obj.name = name
}

func (obj *person) GetName() string{
	return obj.name
}

func (obj *person) String() string {
	return fmt.Sprintf("{name: %v, age: %v}", obj.name, obj.age)
}