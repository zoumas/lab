package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan struct{})

	go func() {
		for {
			j, ok := <-jobs
			if !ok {
				fmt.Println("received all jobs")
				done <- struct{}{}
				return
			}

			fmt.Println(j, "received job")
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	fmt.Println("sent all jobs")
	close(jobs)

	<-done

	_, ok := <-jobs
	fmt.Println("received more jobs:", ok)
}
