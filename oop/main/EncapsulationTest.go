

package main

import (
	"fmt"
	domain "oop/domain"
)


func main(){

	var p = domain.NewPerson("Li")
	p.SetName("hU")
	fmt.Printf("obj: %v \n", p)
	fmt.Printf("p.GetName(): %v \n", p.GetName())


} 