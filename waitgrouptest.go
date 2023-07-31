package main

import (
	"sync"
	"time"
)

func waitGroupTest() {
	var wg sync.WaitGroup
	/*
		waitgroup本质就是维护了一个计数器，add就是往里面
		增加计数器
	*/
	wg.Add(3)

	go func() {
		/*
			waitgroup的Done就是减少计数器
		*/
		defer wg.Done()
		time.Sleep(1000 * time.Millisecond)
	}()

	go func() {
		defer wg.Done()
		time.Sleep(1500 * time.Millisecond)
	}()

	go func() {
		defer wg.Done()
		time.Sleep(2000 * time.Millisecond)
	}()

	/*
		wg中的计数器归0之后，这个wg.Wait就返回了
	*/
	wg.Wait()
	println("3个waitgroup都结束了")
}
