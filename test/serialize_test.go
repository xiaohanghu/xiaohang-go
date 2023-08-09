package test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

func Store(obj any, filePath string) error {
	data, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	fmt.Printf("data: %v\n", string(data))
	err = os.WriteFile(filePath, data, 0666)
	//file , err1 := os.OpenFile("example.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	return err
}

func TestFileMode(t *testing.T) {
	fmt.Println(os.FileMode(0666))
	//fmt.Println(os.FileMode("-rw-rw-rw-"))
	//	[-][rw-][rw-][rw-]
	//  [文件类型][拥有者权限][群组使用者权限][其他使用者权限]
	// r: 可读, w: 可写, x: 可执行
	//	6:	可以读、可以写
	//	7:	可以读、可以写、可以执行

	//fmt.Printf("%T, %v\n", os.ModeDir, os.ModeDir)
	//fmt.Printf("%T, %v\n", os.ModeAppend, os.ModeAppend)
	//fmt.Printf("%T, %v\n", os.ModeExclusive, os.ModeExclusive)
	//fmt.Printf("%T, %v\n", os.ModeTemporary, os.ModeTemporary)
	//fmt.Printf("%T, %v\n", os.ModeSymlink, os.ModeSymlink)
	//fmt.Printf("%T, %v\n", os.ModeDevice, os.ModeDevice)
	//fmt.Printf("%T, %v\n", os.ModeNamedPipe, os.ModeNamedPipe)
	//fmt.Printf("%T, %v\n", os.ModeSocket, os.ModeSocket)
	//fmt.Printf("%T, %v\n", os.ModeSetuid, os.ModeSetuid)
	//fmt.Printf("%T, %v\n", os.ModeSetgid, os.ModeSetgid)
	//fmt.Printf("%T, %v\n", os.ModeCharDevice, os.ModeCharDevice)
	//fmt.Printf("%T, %v\n", os.ModeSticky, os.ModeSticky)
	//fmt.Printf("%T, %v\n", os.ModeIrregular, os.ModeIrregular)
	//fmt.Printf("%T, %v\n", os.ModeType, os.ModeType)
	//fmt.Printf("%T, %v\n", os.ModePerm, os.ModePerm)

	dir := "./data"
	os.Mkdir(dir, 0666)

	//fileInof, _ := os.Stat("./monster.json")
	fileInof, _ := os.Stat(dir)
	fmt.Printf("fileInof: %o \n", fileInof.Mode())
}

func Read(v any, file string) error {
	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, v)
	return err
}

func TestSaveToFile(t *testing.T) {
	m := &Monster{
		Name:  "终结者T0",
		Age:   8,
		Skill: "Run",
	}

	filePath := "./monster.json"
	Store(m, filePath)
}

func TestReadFromFile(t *testing.T) {
	m := &Monster{}
	filePath := "./monster.json"
	Read(m, filePath)
	fmt.Printf("%v \n", m)
}
