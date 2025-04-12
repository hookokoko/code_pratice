package dui

import (
	"fmt"
	"testing"
)

func heapifyMax(arr []int, index int) {
	// index 是作为父节点的索引
	left := 2*index + 1
	right := 2*index + 2
	if left >= len(arr) && right >= len(arr) {
		return
	}
	var largest int
	if left < len(arr) && arr[left] > arr[index] {
		largest = left
	} else {
		largest = index
	}

	if right < len(arr) && arr[right] > arr[largest] {
		largest = right
	}

	if largest != index {
		arr[index], arr[largest] = arr[largest], arr[index]
		heapifyMax(arr, largest)
	}
}

func buildMaxHeap(arr []int) {
	if len(arr) <= 1 {
		return
	}
	for i := len(arr) / 2; i >= 0; i-- {
		heapifyMax(arr, i)
	}
}

func addToMaxHeap(arr *[]int, val int) {
	*arr = append(*arr, val)
	if len(*arr) == 1 {
		return
	}
	buildMaxHeap(*arr)
}

func popMaxHeap(arr *[]int) int {
	largest := (*arr)[0]
	// arr = arr[1:]
	// 其实这里可以直接buildMaxHeap了，但是考虑堆这种数据结构的实现，还是按照经典的方式来实现
	// 即最后一个元素放到root上
	(*arr)[0] = (*arr)[len(*arr)-1]
	*arr = (*arr)[:len(*arr)-1]
	buildMaxHeap(*arr)
	return largest
}

func addToMaxHeap1(arr *[]int, val int) {
	*arr = append(*arr, val)
	if len(*arr) == 1 {
		return
	}
	i := len(*arr) - 1
	for i >= 0 && (*arr)[i/2] < val {
		(*arr)[i/2], (*arr)[i] = (*arr)[i], (*arr)[i/2]
		i = i / 2
	}
}

func popMaxHeap1(arr *[]int) int {
	largest := (*arr)[0]
	// arr = arr[1:]
	// 其实这里可以直接buildMaxHeap了，但是考虑堆这种数据结构的实现，还是按照经典的方式来实现
	// 即最后一个元素放到root上
	(*arr)[0] = (*arr)[len(*arr)-1]
	*arr = (*arr)[:len(*arr)-1]
	heapifyMax(*arr, 0)
	return largest
}

func TestMaxHeap(t *testing.T) {
	arr := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	buildMaxHeap(arr)
	fmt.Println(arr)
	large := popMaxHeap(&arr)
	fmt.Println(large, arr)
}

func heapifyMin(arr []int, index int) {
	// index 是作为父节点的索引
	left := 2*index + 1
	right := 2*index + 2
	if left >= len(arr) && right >= len(arr) {
		return
	}
	var smallest int
	if left < len(arr) && arr[left] < arr[index] {
		smallest = left
	} else {
		smallest = index
	}

	if right < len(arr) && arr[right] < arr[smallest] {
		smallest = right
	}

	if smallest != index {
		arr[index], arr[smallest] = arr[smallest], arr[index]
		heapifyMin(arr, smallest)
	}
}

func buildMinHeap(arr []int) {
	if len(arr) <= 1 {
		return
	}
	for i := len(arr) / 2; i >= 0; i-- {
		heapifyMin(arr, i)
	}
}

func addToMinHeap(arr *[]int, val int) {
	*arr = append(*arr, val)
	if len(*arr) == 1 {
		return
	}
	// 没必要每次add都buildHeap，这样时间复杂度就是nlogn了
	buildMinHeap(*arr)
}

func addToMinHeap1(arr *[]int, val int) {
	*arr = append(*arr, val)
	if len(*arr) == 1 {
		return
	}
	i := len(*arr) - 1
	for i >= 0 && (*arr)[i/2] > val {
		(*arr)[i/2], (*arr)[i] = (*arr)[i], (*arr)[i/2]
		i = i / 2
	}
}

func TestAddToMinHeap1(t *testing.T) {
	arr := []int{1, 4, 5, 7, 9}
	addToMinHeap1(&arr, 3)
	fmt.Println(arr)
}

func popMinHeap(arr *[]int) int {
	smallest := (*arr)[0]
	// arr = arr[1:]
	// 其实这里可以直接buildMaxHeap了，但是考虑堆这种数据结构的实现，还是按照经典的方式来实现
	// 即最后一个元素放到root上
	(*arr)[0] = (*arr)[len(*arr)-1]
	*arr = (*arr)[:len(*arr)-1]
	buildMinHeap(*arr)
	return smallest
}

func popMinHead1(arr *[]int) int {
	smallest := (*arr)[0]
	(*arr)[0] = (*arr)[len(*arr)-1]
	*arr = (*arr)[:len(*arr)-1]
	heapifyMin(*arr, 0)
	return smallest
}

func TestPopMinHeap1(t *testing.T) {
	arr := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	buildMinHeap(arr)
	fmt.Println(arr)
	popMinHead1(&arr)
	fmt.Println(arr)
}

func TestMinHeap(t *testing.T) {
	arr := []int{}
	for _, val := range []int{-1, -2, -3, -4, -5} {
		addToMinHeap(&arr, val)
		fmt.Println(arr)
	}
}

var minHeap = make([]int, 0, 16)
var maxHeap = make([]int, 0, 16)

func AddNumber(num int) {
	l := len(minHeap)
	r := len(maxHeap)
	if l == r {
		if r == 0 || num < minHeap[0] {
			addToMaxHeap(&maxHeap, num)
		} else {
			val := popMinHeap(&minHeap)
			addToMaxHeap(&maxHeap, val)
			addToMinHeap(&minHeap, num)
		}
	} else {
		// l == 0 || 不要加这个！！！
		if num > maxHeap[0] {
			addToMinHeap(&minHeap, num)
		} else {
			val := popMaxHeap(&maxHeap)
			addToMinHeap(&minHeap, val)
			addToMaxHeap(&maxHeap, num)
		}
	}
}

// 超时了
func FindMedian() int {
	l := len(minHeap)
	r := len(maxHeap)
	if l == r {
		return (minHeap[0] + maxHeap[0]) / 2
	} else {
		return maxHeap[0]
	}
}

func TestFind(t *testing.T) {
	arr := []int{
		78,
		14,
		50,
		20,
		13,
		9,
		25,
		8,
		13,
		37,
		29,
		33,
		55,
		52,
		6,
		17,
		65,
		23,
		74,
		43,
		5,
		29,
		29,
		72,
		7,
		13,
		56,
		21,
		31,
		66,
		69,
		69,
		74,
		12,
		77,
		23,
		10,
		6,
		27,
		63,
		77,
		21,
		40,
		10,
		19,
		59,
		35,
		40,
		44,
		4,
		15,
		29,
		63,
		27,
		46,
		56,
		0,
		60,
		72,
		35,
		54,
		50,
		14,
		29,
		62,
		24,
		18,
		79,
		16,
		19,
		8,
		77,
		10,
		21,
		66,
		42,
		76,
		14,
		58,
		20,
		0,
	}
	for _, num := range arr {
		AddNumber(num)
		fmt.Println(FindMedian())
	}
}
