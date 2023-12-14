package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type ReadOp struct {
	Key      int
	Response chan int
}

type WriteOp struct {
	Key      int
	Value    int
	Response chan bool
}

func main() {
	var readOps uint64
	var writeOps uint64

	reads := make(chan ReadOp)
	writes := make(chan WriteOp)

	go func() {
		state := make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.Response <- state[read.Key]
			case write := <-writes:
				state[write.Key] = write.Value
				write.Response <- true
			}
		}
	}()

	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := ReadOp{
					Key:      rand.Intn(5),
					Response: make(chan int),
				}

				reads <- read
				<-read.Response
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := WriteOp{
					Key:      rand.Intn(5),
					Value:    rand.Intn(100),
					Response: make(chan bool),
				}

				writes <- write
				<-write.Response
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)
	readOpsFinal := atomic.LoadUint64(&readOps)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("readOps:", readOpsFinal)
	fmt.Println("writeOps:", writeOpsFinal)
}
