package sort

func partition_my(nums []int, left, right int) {
	pivot := nums[left]
	for left < right {
		for pivot < nums[right] && left < right { // left和right的边界问题搞错
			right--
		}
		nums[left] = nums[right]
		//left++ // 交换之后不用再特意执行加一操作，这是代码和演示的区别
		for pivot > nums[left] && left < right {
			left++
		}
		nums[right] = nums[left]
		//right--
	}
	nums[left] = pivot
}

func quickSort(arr []int, left, right int) {
	if left < right {
		pivot := partition(arr, left, right)
		quickSort(arr, left, pivot-1)
		quickSort(arr, pivot+1, right)
	}
}

func partition(arr []int, left, right int) int {
	pivot := arr[right]
	i := left - 1
	for j := left; j < right; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[right] = arr[right], arr[i+1]
	return i + 1
}
