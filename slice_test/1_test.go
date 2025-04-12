package slice_test

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func PrintSliceStruct(s *[]int) {
	// 代码 将slice 转换成 reflect.SliceHeader
	ss := (*reflect.SliceHeader)(unsafe.Pointer(s))

	// 查看slice的结构
	fmt.Printf("slice struct: %+v, slice is %v\n", ss, s)
}

//func test(s []int) {
//	PrintSliceStruct(&s)
//}
//
//func Test_1(t *testing.T) {
//	s := make([]int, 5, 10)
//	PrintSliceStruct(&s)
//	test(s)
//}
//
//// 底层数组不变
//func case1(s []int) {
//	s[1] = 1
//	PrintSliceStruct(&s)
//}
//
//// 底层数组变化
//func case2(s []int) {
//	s = append(s, 0)
//	s[1] = 1
//	PrintSliceStruct(&s)
//}
//
//func Test_2(t *testing.T) {
//	s := make([]int, 5)
//	case1(s)
//	//case2(s)
//	PrintSliceStruct(&s)
//}

// 截取0号元素以后的元素
func case1(s []int) {
	s = s[1:]
	PrintSliceStruct(&s)
}

// 截取[1, 2]区间元素
func case2(s []int) {
	s = s[1:3]
	PrintSliceStruct(&s)
}

// 截取[len(s)-1, )区间元素
func case3(s []int) {
	s = s[len(s)-1:]
	PrintSliceStruct(&s)
}

// 截取获得新切片
func case4(s []int) {
	s1 := s[2:]
	PrintSliceStruct(&s1)
}

func Test_3(t *testing.T) {
	s := make([]int, 5)
	PrintSliceStruct(&s)
	case1(s)
	case2(s)
	case3(s)
	case4(s)
	PrintSliceStruct(&s)
}

func Test_4(t *testing.T) {
	s := []int{0, 1, 2, 3, 4}

	_ = s[4]
	PrintSliceStruct(&s)
	// 删除第1个元素（从0开始计数）
	// [0, 1) + [2, len(s))
	s1 := append(s[:1], s[2:]...)
	{
		// 拷贝元素
		// 0, 1, 2, 3, 4
		// 0, 2, 3, 4, 4
	}

	PrintSliceStruct(&s1)
	PrintSliceStruct(&s)

	//访问原切片
	_ = s[4]
	//访问从原切片中删除了一个元素的切片
	//_ = s1[4]
}

func Test_5(t *testing.T) {
	//func case1() {
	s1 := make([]int, 3, 3)
	s1 = append(s1, 1)
	PrintSliceStruct(&s1)

	//func case2() {
	ss1 := make([]int, 3, 4)
	ss2 := append(ss1, 1)

	PrintSliceStruct(&ss1)
	PrintSliceStruct(&ss2)

	//func case3() {
	sss1 := make([]int, 3, 3)
	sss2 := append(sss1, 1)

	PrintSliceStruct(&sss1)
	PrintSliceStruct(&sss2)
}

func Test_6(t *testing.T) {
	doAppend := func(s []int) {
		s = append(s, 1)
		printLengthAndCapacity(s)
	}
	s := make([]int, 8, 8)
	doAppend(s[:4])
	printLengthAndCapacity(s)
	fmt.Printf("\n\n")
	doAppend(s)
	printLengthAndCapacity(s)
}

func printLengthAndCapacity(s []int) {
	fmt.Println(s)
	fmt.Printf("len=%d cap=%d \n", len(s), cap(s))
}
