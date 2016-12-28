package huffman

import "testing"

func buildFreshQueue() *RunePriorityQueue {
	return NewRunePriorityQueue(&PriorityQueue{
		&Item{Value: 'a', Priority: .4},
		&Item{Value: 'q', Priority: .6},
		&Item{Value: 'd', Priority: .9}})
}

var pushAndPopTests = []struct {
	beforeQueue *RunePriorityQueue
	pushedItem  *Item
	poppedItem  *Item
}{
	{buildFreshQueue(), &Item{Value: 'F', Priority: 1}, &Item{Value: 'F', Priority: 1}},
	{buildFreshQueue(), &Item{Value: 'F', Priority: .3}, &Item{Value: 'd', Priority: 0.9}},
	{buildFreshQueue(), &Item{Value: 'F', Priority: .7}, &Item{Value: 'd', Priority: .90}},
}

func TestPushAndPop(t *testing.T) {
	for _, test := range pushAndPopTests {
		test.beforeQueue.Push(test.pushedItem)
		topItem := test.beforeQueue.Pop()
		if !topItem.Equals(*test.poppedItem) {
			t.Fatalf("Expected %s but got %s", test.poppedItem.String(), topItem.String())
		}
	}

}
