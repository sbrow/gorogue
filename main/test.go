package main

import (
	"fmt"
	"sync"
)

var p []chan int
var wg sync.WaitGroup

func Conn(ch chan int, priority int, value string) {
	ch <- priority
	<-ch
	fmt.Println(value)
	ch <- 0
}

func main() {
	p = make([]chan int, 5)
	wg.Add(1)
	conns := 5
	for i := 1; i <= conns; i++ {
		c := make(chan int)
		go Conn(c, i, "A"+fmt.Sprint(i))
		p[i-1] = c
	}
	go srv()
	wg.Wait()
}

func srv() {
	defer wg.Done()
	pp := make([]chan int, 5)
	for _, ch := range p {
		pp[5-<-ch] = ch
	}
	for _, ch := range pp {
		ch <- 1
		<-ch
	}
}
