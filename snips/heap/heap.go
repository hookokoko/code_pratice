package heap

type SmallHeap struct {
	Value [][2]int
}

func (sh *SmallHeap) Less(i, j int) bool {
	// 大顶堆就是这样了，其他不用变
	//return sh.Value[i][1] 》 sh.Value[j][1]
	return sh.Value[i][1] < sh.Value[j][1]
}

func (sh *SmallHeap) Pop() interface{} {
	old := (*sh).Value
	n := len(old)
	x := old[n-1]
	(*sh).Value = old[0 : n-1]
	return x
}

func (sh *SmallHeap) Swap(i, j int) {
	sh.Value[i], sh.Value[j] = sh.Value[j], sh.Value[i]
}

func (sh *SmallHeap) Len() int {
	return len(sh.Value)
}

func (sh *SmallHeap) Push(x interface{}) {
	(*sh).Value = append((*sh).Value, x.([2]int))
}
