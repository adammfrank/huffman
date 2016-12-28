package huffman

import (
	"fmt"
)

// Item is a pairing of a rune and its probability in the string
type Item struct {
	Value    interface{}
	Priority float64
	index    int
}

func (this Item) Equals(other Item) bool {
	return this.Value == other.Value
}

func (this Item) String() string {
	return fmt.Sprintf("{Value: %c, Priority: %f, index: %d}", this.Value, this.Priority, this.index)
}
