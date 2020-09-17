// Package unionfind implements a Union-Find (a.k.a. Disjoint Set) data structure
package unionfind

// Node represents a tree node of the disjoint-set forest.
type Node struct {
	Parent *Node
	// Rank is an upper bound on the height of the node, which is the #(edges)
	// of path from this node to its farthest leaf
	Rank int
	Val  interface{}
}

// FindSet queries the root for a given node, and make each node on the find
// path point directly to the root.
func FindSet(x *Node) *Node {
	if x.Parent != x { // x isn't the root
		x.Parent = FindSet(x.Parent)
	}

	return x.Parent
}

// Link the forest of smaller rank to the higher one. In case of tie, link the
// first forest to the second one.
func Link(x, y *Node) *Node {
	if x.Rank > y.Rank {
		y.Parent = x
		return x
	}

	x.Parent = y
	if x.Rank == y.Rank {
		y.Rank++
	}

	return y
}

// MakeSet constructs a plain forest of size 1.
func MakeSet(v interface{}) *Node {
	x := &Node{Val: v}
	x.Parent = x

	return x
}

// Union merges two forest according to "union by rank" heuristic rule.
func Union(x, y *Node) *Node {
	return Link(FindSet(x), FindSet(y))
}
