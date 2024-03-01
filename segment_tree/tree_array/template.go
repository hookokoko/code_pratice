package tree_array

type sumFunc func(val ...any)

var (
	sumArr []int
	oriArr []int
)

func lowBit(x int) int {
	return x & -x
}

func query(x int) int {
	// x从1开始
	res := 0
	for x > 0 {
		res += sumArr[x]
		x = x - lowBit(x)
	}
	return res
}

func add(x int, val int) {
	// x从1开始
	for x < len(sumArr) {
		sumArr[x] += val
		x += lowBit(x)
	}
}

func initSum() {
	oriArr = []int{8, 6, 1, 4, 5, 5, 1, 1, 3, 2, 1, 4, 9, 0, 7, 4}
	sumArr = make([]int, len(oriArr)+1)
	for i := 0; i < len(oriArr); i++ {
		add(i+1, oriArr[i])
	}
}
