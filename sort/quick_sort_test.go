package sort

import (
	"fmt"
	"testing"
)

func Test_quick(t *testing.T) {
	a := []int{19, 97, 9, 17, 1, 8}
	//partition(a, 0, 5) // [1 8 9 17 19 97]
	partition_my(a, 0, 5)
	fmt.Println(a)
}

//output
//arr := []int{5, 2, 6, 3, 1, 4}
//fmt.Println("Before sorting:", arr)
//quickSort(arr, 0, len(arr)-1)
//fmt.Println("After sorting:", arr)
