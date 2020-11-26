package main

import (
	"fmt"
	"time"

	"github.com/gothello/workpool/pool"
)

type task int

func (t task) Run() {
	fmt.Println("doing something")
	time.Sleep(time.Second)
}

func main() {

	p := pool.New(10)

	for i := 0; i < 100; i++ {
		worker := <-p.Workers
		worker <- task(i)

	}
	p.Stop()

}
