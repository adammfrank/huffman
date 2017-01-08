package main

import (
	"fmt"

	"github.com/adammfrank/huffman/huffman"
)

func main() {
	const input = "AABBBCDDEDDDE"
	probs := huffman.CalculateProbabilities(input)
	rpq := huffman.NewRunePriorityQueue(nil)
	for k, v := range probs {
		item := new(huffman.Item)
		item.Value = k
		item.Priority = v
		rpq.Push(item)
	}

	for rpq.Len() > 1 {
		fmt.Println(rpq.Len())
		item1 := rpq.Pop()
		item2 := rpq.Pop()

		parent := new(huffman.Item)
		parent.LeftChild = item1
		parent.RightChild = item2
		parent.Priority = item1.Priority + item2.Priority
		parent.Value = nil
		rpq.Push(parent)
	}

	fmt.Println(rpq.Pop())

}
