package my_moudle

import "fmt"

// func say(s string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(1 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

// 通道
// ch <- v    // 把 v 发送到通道 ch
// v := <-ch  // 从 ch 接收数据
//            // 并把值赋给 v
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c) // 关闭通道
}

func Do_go() {
	// // 并发
	// d0 := time.Now()
	// go say("world")
	// say("hello")
	// d1 := time.Now()
	// fmt.Println(d1.Sub(d0))

	// 通道学习
	// 无缓存
	// s := []int{7, 2, 8, -9, 4, 0}

	// c := make(chan int)
	// go sum(s[:len(s)/2], c)
	// go sum(s[len(s)/2:], c)
	// x, y := <-c, <-c // 从通道 c 中接收

	// fmt.Println(x, y, x+y)

	// // 缓冲通道
	// ch := make(chan int, 2)

	// // 因为 ch 是带缓冲的通道，我们可以同时发送两个数据
	// // 而不用立刻需要去同步读取数据
	// ch <- 1
	// ch <- 2

	// t_v := <-ch
	// // 获取这两个数据
	// fmt.Println(t_v)
	// fmt.Println(&t_v)
	// fmt.Println(<-ch)

	// 关闭通道

	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	// range 函数遍历每个从通道接收到的数据，因为 c 在发送完 10 个
	// 数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据
	// 之后就结束了。如果上面的 c 通道不关闭，那么 range 函数就不
	// 会结束，从而在接收第 11 个数据的时候就阻塞了。
	for i := range c {
		fmt.Println(i)
	}
}
