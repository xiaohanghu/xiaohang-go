package test

import (
	"fmt"
	"testing"
)

// https://pkg.go.dev/testing@go1.20.7
// 文件名必须以 _test.go 结束
// 方法名必须以 Test 开头
func TestHello(t *testing.T) {
	fmt.Println("Hello test")
	t.Errorf("test error %v", 123)
}
