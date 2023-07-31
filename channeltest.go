package main

import "fmt"

/*
go routine可以认为是一种轻量级的线程函数，与之配对使用的就是channel。
当go routine往一个channel里写，且channel满了，则会导致等待。
如果go routine从一个channel里面读取，且channel空，则也会等待。
main函数本身也是一个goroutine函数
*/

func visitChannel() {
	//创建一个缓冲为10的int的channel
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		//这里加上了v这个参数来构成闭包，这样每个go
		//函数capture的值都是不一样的
		go func(v int) {
			ch <- v
		}(i)
	}

	for j := 0; j < 10; j++ {
		fmt.Println(<-ch)
	}

}

func channelDL() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		v := 1
		//这里的ch1没有buffer，所以这行写入将导致等待
		ch1 <- v
		v2 := <-ch2
		fmt.Println(v, v2)
	}()

	v := 2
	//这里的ch2也没有buffer，所以这行写入将导致等待
	//这样，两个goroutein就锁死在这里了
	ch2 <- v
	v2 := <-ch1
	fmt.Println(v, v2)
}

func channelSwitch() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		v := 1
		//这里的ch1没有buffer，所以这行写入将导致等待
		ch1 <- v
		v2 := <-ch2
		fmt.Println(v, v2)
	}()

	v := 2
	var v2 int
	//虽然上面的goroutine依然等待，但是下面的goroutine使用了
	//select，这时候两个case中的任意一个可以执行，就执行下去了
	//这个有点像当年windows中的waitformultipleobjects
	select {
	case ch2 <- v:
	case v2 = <-ch1:

	}
	fmt.Println(v, v2)
}
