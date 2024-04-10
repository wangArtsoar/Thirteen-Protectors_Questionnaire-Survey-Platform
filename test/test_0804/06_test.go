package test_0804

import (
	"fmt"
	"testing"
)

var s = []int{1, 2, 34, 35}
var v = 3

// TestClaudeAnswer
func TestClaudeAnswer(t *testing.T) {
	// 切片s的底层数组
	array := s[:cap(s)]
	// 追加新元素
	array = append(array, v)
	fmt.Println(array[:len(s)+1])
}

// TestFirstAnswer
func TestFirstAnswer(t *testing.T) {
	// 创建一个切片
	mySlice := []string{"hello", "world", "this", "is", "a", "test"}

	// 向切片中添加元素
	newElement := "append me!"
	appendedSlice := append(mySlice, newElement)

	// 检查数组是否已更改
	fmt.Println("Original array before append:", mySlice)
	fmt.Println("Appended array after append:", appendedSlice)
}

// ignore
func appendInPlace(s []int, v int) []int {
	// 切片s的底层数组
	array := s[:cap(s)]
	// 追加新元素
	array = append(array, v)
	// 返回更新后的切片
	return array[:len(s)+1]
}
