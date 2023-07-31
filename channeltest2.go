package main

import (
	"fmt"
	"time"
)

func channelSwitch2() {
	ch1 := make(chan int, 10)
	stop := make(chan struct{})

	/*
		第一个go-routine，在内部循环10次，把0-9的数字放入
		这个channel ch1中
	*/
	go func(val int) {
		for i := 0; i < val; i++ {
			ch1 <- i
		}
	}(10)

	/*
		第二个go-routine，在休息一段时间后，在top channel
		中写入一个struct{}对象
	*/
	go func() {
		time.Sleep(1000 * time.Millisecond)
		stop <- struct{}{}
	}()

	var v2 int
	/*
		第三个go-routine（main这个），进入无限循环，内部使用
		了一个select命令（相当于waitformultipleobjects），
		如果stop channel中有数据，就直接结束这个routine，如果
		是ch1中有数据，则输出并重新进入这个循环
	*/
	for i := 0; ; i++ {
		select {
		case _ = <-stop:
			fmt.Println("ended")
			return
		case v2 = <-ch1:
			fmt.Print(v2)
			fmt.Println("  v2")
			continue
		}

	}

}
