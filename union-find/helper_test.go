package uf_test

import "github.com/sammy00/uf"

func makePath(ell int) (*uf.Node, *uf.Node) {
	leaf := new(uf.Node)
	leaf.Parent = leaf

	x, root := leaf, leaf

	for ; ell > 0; ell-- {
		root, x = new(uf.Node), root
		x.Parent, root.Parent, root.Rank = root, root, x.Rank+1
	}

	return leaf, root
}
