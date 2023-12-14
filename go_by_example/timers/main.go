package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.NewTimer(2 * time.Second)

	<-t1.C
	fmt.Println("Timer 1 fired")

	t2 := time.NewTimer(time.Second)
	go func() {
		<-t2.C
		fmt.Println("Timer 2 fired")
	}()

	stop2 := t2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
	time.Sleep(2 * time.Second)
}
