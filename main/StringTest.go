package main

import (
	"fmt"
	// "unsafe"
	"unicode/utf8"
	"strings"
)

/*
字符使用Unicode字符集和UTF-8编码
一个utf8数字占1个字节
一个utf8英文字母占1个字节
少数是汉字每个占用3个字节，多数占用4个字节。
*/ 

func main() {
	var c1 byte = 'y'
	fmt.Println("c1:", c1)
	fmt.Printf("c1：%c", c1)
	fmt.Println()

	text1,text2,text3 := "abc", "啊", "测试"
	// text1[0]='f' 字符串是不可变的， 不能修改字符串中的值
	fmt.Printf("text1占 %d bits, text2占 %d bits", len(text1), len(text2)) // len(str)按字节统计字符串长度 
	fmt.Println()
	
	fmt.Printf("text3长度为：%d", utf8.RuneCountInString(text3))
	fmt.Println()

	s1 := `
func main() {
}
	`
	fmt.Println("s1:", s1)

	s2 := "abc" + // 换行的话+保留在最后面
	"sdfs"
	fmt.Println("s2:", s2)


	str := "hahatest你好"

	// 字符串遍历
	fmt.Println("用range遍历：")
	for i, v := range str { 
		fmt.Printf("str[%d]=%c %T \n", i, v, v) // 注意：i 是按byte的下标，并不是字符的下标
	}

	fmt.Println("------------------------")
	fmt.Println("用切片遍历：")
	r := []rune(str)
	for i :=0; i < len(r); i++ {
		v := r[i]
		fmt.Printf("str[%d]=%c %T \n", i, v, v)
	}

	count := strings.Count("golanggo", "go") // 统计go出现的次数, 大小写敏感
	fmt.Printf("count:%v \n", count)

	fmt.Printf("string equal ignore case: %v \n", strings.EqualFold("go", "Go")) // == 比较区分大小写 

	fmt.Printf("strings.Index: %v \n", strings.Index("golanggago", "ga"))
 
	// 将s中前n个不重叠old子串都替换为new的新字符串, n<0则全部替换
	fmt.Printf("strings.Replace: %v \n", strings.Replace("golanggago", "go", "#", -1))

	fmt.Printf("strings.Split: %v \n", strings.Split("go-python-java", "-"))

	fmt.Printf("strings.ToUpper: %v \n", strings.ToUpper("go-python-java"))
	fmt.Printf("strings.ToLower: %v \n", strings.ToLower("GO-PYTHON-JAVA"))


	temp := `strings.TrimSpace: "%v"
`
	fmt.Printf(temp, strings.TrimSpace("  go python java  "))

	fmt.Printf("strings.TrimRight: \"%v\" \n", strings.TrimRight("~go python java~~", "~"))
	// strings.TrimLeft

	fmt.Printf("strings.HasPrefix: \"%v\" \n", strings.HasPrefix("~go python java~~", "~go"))
	// strings.HasSuffix() 

	s5 := "adfjwoerwero" // string 是不可变的，需要转化成数组操作
	var s5Byte = []byte(s5)
	s5Byte[1] = 'b'
	s5 = string(s5Byte)
	fmt.Printf("s5Byte: \"%v\" \n", s5Byte)
	fmt.Printf("s5: \"%v\" \n", s5)

	var s5Rune = []rune(s5) // Rune是以字符为单位的数组
	s5Rune[1] = '我'
	s5 = string(s5Rune)
	fmt.Printf("s5Rune: \"%v\" \n", s5Rune)
	fmt.Printf("s5: \"%v\" \n", s5)

}
