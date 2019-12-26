package main

import (
	"fmt"
	"sync"
	"time"
)

var ch = make(chan int, 2)

func main() {

	// runtime.GOMAXPROCS(4)
	// fmt.Println("NUmer of cores:", runtime.NumCPU())
	// ch := make(chan int, 1)
	// go func() {
	// 	fmt.Println("call")
	// 	ch <- 1
	// }()
	// process()
	// <-ch
	// g := gin
	// count := 0

	var wg sync.WaitGroup
	wg.Add(1)

	// for i := 0; i < 3; i++ {
	// go func(a int) {
	// 	for i := 0; i < 3; i++ {
	// count += a

	go func() {
		defer wg.Done()
		ch <- 1
		fmt.Println("OP:")
	}()

	test()

	<-ch
	wg.Wait()

	// ch <- 3
	// ch <- 2
	// fmt.Println("Inc by:%d", i)
	// }
	// }(1)
	// }
	// wg.Wait()
	// for j := 0; j < 3; j++ {

	// }

}

func process() {
	fmt.Println("IN")
	time.Sleep(5 * time.Second)
	fmt.Println("OUT")
}

func test() {
	fmt.Println("hllo")
}

func recoverPanic() {}
