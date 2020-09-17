package unionfind_test

import (
	"testing"

	uf "github.com/sammyne/data-structure-go/union-find"
)

func TestFindSet(t *testing.T) {
	root := new(uf.Node)
	root.Parent = root

	testCases := []struct {
		x      *uf.Node
		expect *uf.Node
	}{
		{
			&uf.Node{
				Parent: &uf.Node{
					Parent: &uf.Node{
						Parent: root,
					},
				},
			},
			root,
		},
		{ // a single node
			root,
			root,
		},
	}

	for i, c := range testCases {
		var path []*uf.Node

		for y := c.x; y != y.Parent; y = y.Parent {
			path = append(path, c.x)
		}

		if parent := uf.FindSet(c.x); parent != c.expect {
			t.Fatalf("#%d invalid parent", i)
		}

		if c.x.Parent != c.expect {
			t.Fatalf("#%d x's parent should be updated", i)
		}

		for _, y := range path {
			if y.Parent != c.expect {
				t.Fatalf("#%d parent of some node isn't updated", i)
			}
		}
	}
}

func TestLink(t *testing.T) {
	a, b := &uf.Node{Rank: 123}, &uf.Node{Rank: 128}
	c, d := &uf.Node{Rank: 123}, &uf.Node{Rank: 123}
	e, f := &uf.Node{Rank: 123}, &uf.Node{Rank: 120}

	type expect struct {
		root *uf.Node
		rank int
	}
	testCases := []struct {
		x, y   *uf.Node
		expect expect
	}{
		{a, b, expect{b, b.Rank}},
		{c, d, expect{d, d.Rank + 1}},
		{e, f, expect{e, e.Rank}},
	}

	for i, c := range testCases {
		got := uf.Link(c.x, c.y)

		if got != c.expect.root {
			t.Fatalf("#%d invalid new root: got %v, expect %v", i,
				got, c.expect.root)
		}

		if got.Rank != c.expect.rank {
			t.Fatalf("#%d invalid new rank: got %d, expect %d", i, got.Rank,
				c.expect.rank)
		}
	}
}

func TestUnion(t *testing.T) {
	a, _ := makePath(2)
	c, d := makePath(3)

	e, _ := makePath(2)
	g, h := makePath(2)

	i, j := makePath(3)
	k, _ := makePath(2)

	testCases := []struct {
		x, y   *uf.Node
		expect *uf.Node
	}{
		{a, c, d},
		{e, g, h},
		{i, k, j},
	}

	for i, c := range testCases {
		got := uf.Union(c.x, c.y)
		if got != c.expect {
			t.Fatalf("#%d invalid root: got %v, expect %v", i, got, c.expect)
		}
	}
}
