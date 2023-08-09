package encapsulation

import (
	"fmt"
)

type Animal struct{
	Age int
	weight int
	Name string
}

func (obj *Animal) SetAge(age int) {
	obj.Age = age
}

func (obj *Animal) GetAge() int{
	return obj.Age
}

func (obj *Animal) SetWeight(weight int) {
	obj.weight = weight
}

func (obj *Animal) GetWeight() int{
	return obj.weight
}

func (obj *Animal) String() string {
	return fmt.Sprintf("{Age: %v, weight:%v, Name:%v}", obj.Age, obj.weight, obj.Name)
}