package main

import (
	"fmt"
	"sort"
)

func test(arr []int16){
	arr[0] = 123
}

func main(){

	var arr = [7]int16{4, 2, 6, 1, 9, 33, 87}

	// slice表示一个拥有相同类型元素的可变长度序列
	sl := arr[2:5] //切片 包含1，不包含3. 切片是一个结构体,包含开始地址，切片长度，切片容量
	arr[2] = 222
	test(sl) // slice是引用类型 
	fmt.Printf("type: %T, slice: %v \n", sl, sl)
	fmt.Println("len(sl):", len(sl))
	fmt.Println("slice容量:", cap(sl)) // 容量切片的第一个元素开始数，到其底层数组元素末尾的个数
	fmt.Println("&arr[2] == &sl[0]:", &arr[2] == &sl[0] )

	sl1 := make([]int, 4, 20) //分配内存 make(Type, len, cap)
	// cap的用处：当用append扩展长度时，如果新的长度小于容量，不会更换底层数组，否则，go会新申请一个底层数组，拷贝这边的值过去，把原来的数组丢掉。
	// sl1[5] = 10 // index out of range [5] with length 4
	fmt.Println("sl1:", sl1)

	// var strSlice1 = []string{"Dave", "Tom", "Mary"} // cap是元素的个数
	var strSlice1 = make([]string, 3, 6)
	strSlice1[0] = "Dave"
	strSlice1[1] = "Tom"
	strSlice1[2] = "Mary"

	var strSlice2 = append(strSlice1, "Hang", "Ping") // 会自动扩容, 如果没有扩容就会跟扩容前的slice公用空间
	strSlice1[0] = "TTT"
	fmt.Printf("strSlice1 len: %v, cap: %v, content:%v \n", len(strSlice1), cap(strSlice1), strSlice1)
	fmt.Printf("strSlice2 len: %v, cap: %v, content:%v \n", len(strSlice2), cap(strSlice2), strSlice2)

	strSlice2 = append(strSlice2, strSlice2...)
	fmt.Printf("strSlice2 :%v \n", strSlice2)

	var strSlice3 = make([]string, 2, 6)
	copy(strSlice3, strSlice1) // 超过长度的部分会被丢弃 
	fmt.Printf("strSlice3: %v \n", strSlice3)


	 var str = "wersdflsafjasldfjsdf"
	 var strSl = str[2:9] // 是以byte为单位截取的，不是以字符为单位
	//  strSl[1] = byte('b')
	 fmt.Printf("type:%T, strSl: %v \n", strSl, strSl)

	 var slice1 = []int{4, 2, 6, 1, 9, 33, 87}
	 sort.Ints(slice1)
	 fmt.Printf("slice1: %v \n", slice1)
	
}