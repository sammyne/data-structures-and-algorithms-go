package skiplist

import (
	"testing"
)

func TestNodeString(t *testing.T) {
	n := &Node{
		maxLevel: 456,
		value:    123,
	}

	const expect = `{"MaxLevel":456,"Value":123}`
	if got := n.String(); expect != got {
		t.Fatalf("fail: expect '%s', got '%s'", expect, got)
	}
}
