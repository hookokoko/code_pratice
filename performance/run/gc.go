package main

import (
	"encoding/json"
	"fmt"
	"runtime"
	"sync"
)

// 用于解码数据的临时结构
type QMessage struct {
	ID   uint64       `json:"id,omitempty"`
	Body QMessageBody `json:"body,omitempty"`
}
type QMessageBody struct {
	Field1 string `json:"field_1,omitempty"`
	Field2 int    `json:"field_2,omitempty"`
}

// 常驻于内存的数据结构
type Message struct {
	id     uint64
	field1 string
	field2 int
}

// 内存数据缓存，里面存放了千万级的词条信息
var buffer = make(map[int]Message)

func main() {
	wg := &sync.WaitGroup{}
	q := make(chan string)
	wg.Add(2)
	go producer(q, wg)
	go consumer(q, wg)
	wg.Wait()
	//PrintMemUsage()
}

// 模拟生产者，产生两千万词条数据
func producer(q chan string, wg *sync.WaitGroup) {
	for i := 0; i < 20000000; i++ {
		q <- `{"id":123456, "body":{"field1": "123", "field2": 456}}`
	}
	close(q)
	wg.Done()
}

// 模拟消费者，消费并反序列化词条数据，使用中间临时数据结构进行数据变形，并将最后的结果存储
func consumer(q chan string, wg *sync.WaitGroup) {
	idx := 0
	for data := range q {
		idx++
		qtmp := QMessage{}
		json.Unmarshal([]byte(data), &qtmp)

		tmp := Message{
			id:     qtmp.ID,
			field1: qtmp.Body.Field1,
			field2: qtmp.Body.Field2,
		}
		buffer[idx] = tmp
	}
	wg.Done()
}

// 以下是打印内存监控数据的工具函数，与业务逻辑无关
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v", m.NumGC)
	fmt.Printf("\tAllocObjCnt = %v", m.Mallocs)
	fmt.Printf("\tSTW = %v\n", m.PauseTotalNs)

}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

// =====================================================================================================================
//import (
//	"os"
//	"runtime"
//	"runtime/trace"
//	"sync/atomic"
//)
//
//var stop uint64
//
//// 通过对象 P 的释放状态，来确定 GC 是否已经完成
//func gcfinished() *int {
//	p := 1
//	runtime.SetFinalizer(&p, func(_ *int) {
//		println("gc finished")
//		atomic.StoreUint64(&stop, 1) // 通知停止分配
//	})
//	return &p
//}
//
//func allocate() {
//	// 每次调用分配 0.25MB
//	_ = make([]byte, int((1<<20)*0.25))
//}
//
//func main() {
//	f, _ := os.Create("trace.out")
//	defer f.Close()
//	trace.Start(f)
//	defer trace.Stop()
//
//	gcfinished()
//
//	// 当完成 GC 时停止分配
//	for n := 1; atomic.LoadUint64(&stop) != 1; n++ {
//		println("#allocate: ", n)
//		allocate()
//	}
//	println("terminate")
//}
