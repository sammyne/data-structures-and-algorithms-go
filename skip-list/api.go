package skiplist

import (
	"crypto/rand"
	"errors"
	"fmt"
	"strings"
)

const MaxLevel = 16

var (
	ErrValueNotFound = errors.New("no such value")
	ErrLevelNotFound = errors.New("no such level")
)

type Node struct {
	forwards [MaxLevel]*Node
	maxLevel int // 0-base index of max level
	value    int
}

type SkipList struct {
	maxLevel int // 0-base index of max levels
	lists    Node
}

func (n *Node) MaxLevel() int {
	return n.maxLevel
}

func (n *Node) String() string {
	return fmt.Sprintf(`{"MaxLevel":%d,"Value":%d}`, n.maxLevel, n.value)
}

func (n *Node) Value() int {
	return n.value
}

func (l *SkipList) Delete(v int) {
	toBeUpdated := make([]*Node, l.maxLevel+1)
	for i := range toBeUpdated {
		toBeUpdated[i] = &l.lists
	}

	// at every level, find the largest value smaller than v, these values need updating
	ptr := &l.lists
	for i := l.maxLevel; i >= 0; i-- {
		for ptr.forwards[i] != nil && ptr.forwards[i].value < v {
			ptr = ptr.forwards[i]
		}

		toBeUpdated[i] = ptr
	}

	if ptr.forwards[0] == nil || ptr.forwards[0].value != v {
		return
	}

	// now we have toBeUpdated[i]<v
	for i := l.maxLevel; i >= 0; i-- {
		if ptr.forwards[i] == nil || ptr.forwards[0].value != v {
			continue
		}
		toBeUpdated[i].forwards[i] = toBeUpdated[i].forwards[i].forwards[i]
	}
}

func (l *SkipList) Find(v int) (*Node, error) {
	ptr := &l.lists
	// traversal down to level-0
	for level := l.maxLevel; level >= 0; level-- {
		for ptr.forwards[level] != nil && ptr.forwards[level].value < v {
			ptr = ptr.forwards[level]
		}
	}

	vv := ptr.forwards[0]
	if vv == nil || vv.value != v {
		return nil, ErrValueNotFound
	}

	return vv, nil
}

func (l *SkipList) Insert(v int) {
	// We may check if v already exists. If yes, the following workflow is redundant.
	// Otherwise, just trade off the space of toBeUpdated.

	newNode := &Node{
		value:    v,
		maxLevel: mustRandLevel(),
	}

	toBeUpdated := make([]*Node, newNode.maxLevel+1)
	for i := range toBeUpdated {
		toBeUpdated[i] = &l.lists
	}

	// at every level, find the largest value smaller than v, these values need updating
	ptr := &l.lists
	for i := newNode.maxLevel; i >= 0; i-- {
		for ptr.forwards[i] != nil && ptr.forwards[i].value < v {
			ptr = ptr.forwards[i]
		}

		// value already exists
		if vv := ptr.forwards[i]; vv != nil && vv.value == v {
			return
		}

		toBeUpdated[i] = ptr
	}

	// now we have toBeUpdated[i]<v
	for i := newNode.maxLevel; i >= 0; i-- {
		toBeUpdated[i].forwards[i], newNode.forwards[i] = newNode, toBeUpdated[i].forwards[i]
	}

	if l.maxLevel <= newNode.maxLevel {
		l.maxLevel = newNode.maxLevel
	}
}

func (l *SkipList) LevelString(level int) (string, error) {
	if l.maxLevel < level {
		return "", ErrLevelNotFound
	}

	var out string
	for ptr := &l.lists; ptr.forwards[level] != nil; ptr = ptr.forwards[level] {
		out += fmt.Sprintf("%d ", ptr.forwards[level].value)
	}

	return strings.TrimSuffix(out, " "), nil
}

func (l *SkipList) String() string {
	out, _ := l.LevelString(0)
	return out
}

func NewSkipList() *SkipList {
	return new(SkipList)
}

// why?
func mustRandLevel() int {
	level := 0

	var v [1]byte
	for i := 0; i < MaxLevel; i++ {
		if _, err := rand.Read(v[:]); err != nil {
			panic(err)
		}

		if v[0]%2 == 1 {
			level++
		}
	}

	return level
}
