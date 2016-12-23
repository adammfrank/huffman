package main

import (
	"fmt"

	"github.com/adammfrank/huffman/huffman"
)

func main() {
	const input = "AABBBCDDEDDDE"
	probs := huffman.CalculateProbabilities(input)
	rpq := huffman.NewRunePriorityQueue()
	fmt.Println(probs)
	for k, v := range probs {
		item := new(huffman.Item)
		item.Value = k
		item.Priority = v
		rpq.Push(item)
		fmt.Println(rpq)
	}

	fmt.Println(rpq)

	item1 := &huffman.Item{Value: 'Z', Priority: float64(0.08)}
	rpq.Push(item1)

	fmt.Println(rpq)

	rpq.Pop()

	fmt.Println(rpq)

}
