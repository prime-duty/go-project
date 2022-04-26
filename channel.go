package main

import (
	"fmt"
	"time"
)

type Token struct{}

//循环打印1234
func AlternatelyPrint(total int) {
	//初始化一个channel数组
	channels := make([]chan Token, total)
	for i := 0; i < total; i++ {
		channels[i] = make(chan Token)
	}

	for i := 0; i < total; i++ {
		go func(index int, current chan Token, nextChan chan Token) {
			for {
				<-current
				fmt.Println(index)
				time.Sleep(time.Second)
				nextChan <- Token{}
			}
		}(i+1, channels[i], channels[(i+1)%total])
	}
	channels[0] <- Token{}
	select {}
}

func main() {
	AlternatelyPrint(4)
}
