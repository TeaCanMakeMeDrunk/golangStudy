package goroutineStudy

import (
	"fmt"
	"runtime"
	"time"
)

func Main() {
	//创建一个协程,带参
	go func(a int, b int) {
		fmt.Printf("a=%d, b=%d\n", a, b)
	}(100, 200)

	//不带参
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("====sub_goroutine_", i)
			if i==3{
				runtime.Goexit() //退出
			}
		}
	}()

	fmt.Println("main_goroutine") //相当于一个线程，如果它执行完了，协程一定会退出
	time.Sleep(5*time.Second)
}
