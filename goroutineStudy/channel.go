package goroutineStudy

import "fmt"

func ChannelMain() {
	c := make(chan int, 2) //cap == 2
	go func() {
		for i := 0; i < 5; i++ {
			// 容量大于2时会阻塞等待
			c <- i
			if i == 3 {
				close(c) //关闭一个channel
				break
			}
		}
	}()

	for data := range c {
		fmt.Println("main_goroutine get ", data)
	}
}

func ChannelSelect() {
	c := make(chan int, 2)
	quit := make(chan int) //默认cap == 1

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(<-c) //读取select中写入的数据
		}
		quit <- 0
	}()

	quitChannel(c, quit)

}

func quitChannel(c chan int, quit chan int) {
	x := 0
	for {
		select { //控制多个channel状态
			case c <- x:
				//如果c可写，就会进来
				fmt.Sprintln(" c -< x:", x)
			case <-quit:
				//只要quit chan有数据可读, 就进来
				fmt.Println("quit ..")
				return
			}
	}
}
