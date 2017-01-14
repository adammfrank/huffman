package huffman

import "testing"

var buildTreeTests = []struct {
	input                string
	expectedHeadPriority int
}{
	{"ABCD", 4},
	{"DDCADCCKODD", 11},
	{"KIENFFKKSNKKDCI", 15},
}

func TestBuildItemTreeFromString(t *testing.T) {
	for _, test := range buildTreeTests {
		tree := BuildItemTreeFromString(test.input)
		if tree.Priority != test.expectedHeadPriority {
			t.Fatalf("Expected %d but got %d", test.expectedHeadPriority, tree.Priority)
		}
	}
}
