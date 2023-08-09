package main

import (
	"fmt"
	domain "xiaohang-go/oop/domain"
)

func test(a *domain.Animal) {
	fmt.Printf("animal.Age:%v \n", a.Age)
}

func main() {

	var cat = &domain.Cat{
		// Age : 21, 初始化这里不可以直接给父类字段赋值
		Animal: domain.Animal{
			Age: 21,
			// weight: 31, // 不可访问
			Name: "aaa",
		},
	}
	// var a domain.Animal = cat.(domain.Animal)
	test(&cat.Animal) //go不支持强转，也不支持继承的多态
	fmt.Printf("cat:%v \n", cat)

	cat.Age = 10 // 先在子类里找，找不到再找父类。 而 cat.Animal.Age = 10 则是指定访问父类的属性
	fmt.Printf("cat:%v \n", cat)

	cat.SetAge(11) // cat.Animal.SetAge(11)
	cat.SetWeight(23)
	// cat.SetWeight(3, 4)
	fmt.Printf("cat:%v \n", cat)

	cat.Name = "catName"
	cat.Animal.Name = "animalName"
	fmt.Printf("cat.Name:%v \n", cat.Name)
	fmt.Printf("cat.Animal.Name:%v \n", cat.Animal.Name)

}
