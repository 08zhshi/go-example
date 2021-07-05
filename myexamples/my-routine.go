package main

import (
	"fmt"
	"time"
)

//const (
//	a, b = iota, iota+1
//	c,d
//	e,f
//	g = iota
//)
//type woker struct {
//	Do func()
//}

func main() {
	//var wg sync.WaitGroup
	//ch := make(chan woker, 10)
	//
	//for i := 0; i < 20; i ++ {
	//	go func() {
	//		wg.Add(1)
	//		defer wg.Done()
	//		for job := range ch {
	//			job.Do()
	//		}
	//	}()
	//}
	//
	//for j := 0; j < 100; j ++ {
	//	k := j
	//	w := woker {
	//		Do: func() {
	//			// 输出数据j和k不一样，j大量重复数字，而k没有。
	//			// 原因j是非临时变量，其协程访问其地址指向值，可能导致几个协程访问同一个值。
	//			fmt.Println("j=",j,",k=",k)
	//		},
	//	}
	//	ch <- w
	//}
	//close(ch)
	//wg.Wait()
	//fmt.Println("woker end")
	//// iota值每行递增，同一行iota相等，没有表达式集成前面有的。
	//fmt.Println(a,b,c,d,e,f,g) // 0 1 1 2 2 3 3
	c := make(chan int, 0)
	go func() {
		//select {
		//case i, ok := <- c:
		//	fmt.Println(i, ok)
		//}
		for d := range c {
			fmt.Println(d)
		}
	}()
	close(c)
	time.Sleep(time.Second * 3)
	a := []int{1, 2, 3}
	for i, j := range a {
		fmt.Println(i, j)
	}
}
