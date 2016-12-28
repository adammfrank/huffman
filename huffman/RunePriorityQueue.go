package huffman

import "container/heap"

// RunePriorityQueue wraps a PriorityQueue and turns it into a heap
type RunePriorityQueue struct {
	queue *PriorityQueue
}

// NewRunePriorityQueue constructs a RunePriorityQueue and optionally takes a PriorityQueue
// to use as a base
func NewRunePriorityQueue(pq *PriorityQueue) *RunePriorityQueue {
	newRPQ := new(RunePriorityQueue)
	if pq != nil {
		newRPQ.queue = pq
	} else {
		newRPQ.queue = new(PriorityQueue)
	}
	heap.Init(newRPQ.queue)
	return newRPQ
}

// Push wraps PriorityQueue Push in heap
func (rpq *RunePriorityQueue) Push(item *Item) {
	heap.Push(rpq.queue, item)
}

// Pop wraps PriorityQueue Pop in heap
func (rpq *RunePriorityQueue) Pop() *Item {
	return heap.Pop(rpq.queue).(*Item)
}

func (rpq RunePriorityQueue) String() string {
	result := ""
	for _, v := range *rpq.queue {
		result += v.String()
	}

	return result
}
