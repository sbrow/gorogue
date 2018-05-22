package gorogue

import "container/heap"

// Item is something we manage in a priority queue.
type Item struct {
	Ch       chan string // A communication channel.
	priority int         // Priority of the Item in the queue.
	index    int         // Index of the Item in the heap.
}

func NewItem(ch chan string, priority int) *Item {
	return &Item{Ch: ch, priority: priority}
}

// PriorityQueue implements heap.Interface and holds Items.
type priorityQueue []*Item

func (pq priorityQueue) Len() int { return len(pq) }

// Less returns whether i has a highest priority than j.
func (pq priorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	Item := x.(*Item)
	Item.index = n
	*pq = append(*pq, Item)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	Item := old[n-1]
	Item.index = -1 // for safety
	*pq = old[0 : n-1]
	return Item
}

// update modifies the priority and ch of an Item in the queue.
func (pq *priorityQueue) update(Item *Item, ch chan string, priority int) {
	Item.Ch = ch
	Item.priority = priority
	heap.Fix(pq, Item.index)
}
