package interview

import (
	"fmt"
	"testing"
)

// 使用两个协程交替顺序打印1-10
func print10() {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})

	quit := make(chan struct{})
	i := 0

	go func() {
		// 为了保证这里是协程1和协程2交替顺序打印这里不能用select，用了select之后，case分之就是随机的了
		// 但是，不用select，quit信号的获取会有一个小问题，也就是下面注释提到的，如果不在close之后紧跟一个return，就会多打印一次

		//for {
		//	select {
		//	case <-ch1:
		//		i++
		//		fmt.Println("1---", i)
		//		if i == 10 {
		//			close(quit)
		//		}
		//	case ch2 <- struct{}{}:
		//	case <-quit:
		//		return
		//	}
		//}

		for {
			select {
			case <-quit:
				return
			default:
			}

			<-ch1
			i = i + 1
			fmt.Println("1---", i)
			// 这里有点讲究，之前是没有return，导致协程2还会收到ch2的信号，也就是是会多加打印一次。
			// 同时，还要保证这里的close不能执行2次。close一个已经close的chan会panic，
			// 实际也是的，因为按顺序打印，只可能有一个goroutine进入i==10的分支
			if i == 10 {
				close(quit)
				return
			}
			ch2 <- struct{}{}

		}
	}()

	go func() {
		for {
			select {
			case <-quit:
				return
			default:
			}

			<-ch2
			i = i + 1
			fmt.Println("2---", i)
			if i == 10 {
				close(quit)
				return
			}
			ch1 <- struct{}{}
		}
	}()

	ch1 <- struct{}{}

	<-quit
}

func TestRun(t *testing.T) {
	print10()
}
