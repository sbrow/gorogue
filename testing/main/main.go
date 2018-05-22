package main

import (
	"container/heap"
	"fmt"
	. "github.com/sbrow/gorogue/tests"
	"sync"
)

func Conn(ch chan string, priority int, v string) {
	prequeue <- New(ch, priority)
	<-ch
	fmt.Println(v)
	ch <- "Done!"
}

func srv() {
	defer wg.Done()

	for pq.Len() < 5 {
		heap.Push(&pq, <-prequeue)
	}
	for pq.Len() > 0 {
		i := heap.Pop(&pq).(*Item)
		i.Ch <- "Permission Granted!"
		<-i.Ch
	}
}

var wg sync.WaitGroup
var conns = 5
var pq = PriorityQueue{}
var prequeue = make(chan *Item)
var p = make([]chan string, conns)

func main() {
	heap.Init(&pq)
	wg.Add(1)
	for i := 1; i <= conns; i++ {
		c := make(chan string)
		go Conn(c, i, "A"+fmt.Sprint(i))
		p[i-1] = c
	}
	srv()
	wg.Wait()
}
