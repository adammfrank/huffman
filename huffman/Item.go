package huffman

import (
	"fmt"
)

// Item is a pairing of a rune and its probability in the string
type Item struct {
	Value      interface{}
	Priority   float64
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
	return fmt.Sprintf("{Value: %c, Priority: %f, index: %d}", item.Value, item.Priority, item.index)
}
