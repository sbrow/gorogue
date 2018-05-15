package main

import (
	"container/heap"
	"fmt"
	"sync"
)

// An Item is something we manage in a priority queue.
type Item struct {
	value    interface{} // The value of the item; arbitrary.
	priority int         // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value chan int, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func Conn(ch chan string, priority int, value string) {
	prequeue <- &Item{value: ch, priority: priority}
	<-ch
	fmt.Println(value)
	ch <- "Done!"
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

func srv() {
	defer wg.Done()

	for pq.Len() < 5 {
		heap.Push(&pq, <-prequeue)
	}
	for pq.Len() > 0 {
		i := heap.Pop(&pq).(*Item)
		ch := i.value.(chan string)
		ch <- "Permission Granted!"
		<-ch
	}
}
