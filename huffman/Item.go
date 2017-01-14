package huffman

import "fmt"

// Item is a pairing of a rune and its probability in the string
type Item struct {
	Value      interface{}
	Priority   int
	index      int
	LeftChild  *Item
	RightChild *Item
	Parent     *Item
}

// Equals checks equality of two Items
func (item Item) Equals(other Item) bool {
	return item.Value == other.Value
}

// String prints an Item
func (item Item) String() string {
	return fmt.Sprintf("{Value: %c, Priority: %d, index: %d}", item.Value, item.Priority, item.index)
}

func BuildItemTreeFromString(input string) *Item {
	rpq := NewRunePriorityQueueFromString(input)
	for rpq.Len() > 1 {
		fmt.Println(rpq.Len())
		item1 := rpq.Pop()
		item2 := rpq.Pop()

		parent := new(Item)
		parent.LeftChild = item1
		parent.RightChild = item2
		parent.Priority = item1.Priority + item2.Priority
		parent.Value = nil
		rpq.Push(parent)
	}
	return rpq.Pop()
}
