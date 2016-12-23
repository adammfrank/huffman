package huffman

import "container/heap"

type RunePriorityQueue struct {
	queue *PriorityQueue
}

func NewRunePriorityQueue() *RunePriorityQueue {
	newRPQ := new(RunePriorityQueue)
	newRPQ.queue = new(PriorityQueue)
	heap.Init(newRPQ.queue)
	return newRPQ
}

func (rpq *RunePriorityQueue) Push(item *Item) {
	heap.Push(rpq.queue, item)
}

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
