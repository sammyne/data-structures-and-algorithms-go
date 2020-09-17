package unionfind_test

import (
	"fmt"

	uf "github.com/sammyne/data-structure-go/union-find"
)

func ExampleMakeSet() {
	a, b, c, d := uf.MakeSet(1), uf.MakeSet(1), uf.MakeSet(1), uf.MakeSet(1)

	x := uf.Union(a, b)
	if x == b {
		fmt.Printf("union of a and b produces new root as b of rank=%d\n", b.Rank)
	}

	y := uf.Union(x, c)
	if y == x {
		fmt.Printf("union of x and c produces new root as x of rank=%d\n", x.Rank)
	}

	z := uf.Union(d, y)
	if z == x {
		fmt.Printf("union of y and d produces new root as x of rank=%d\n", x.Rank)
	}

	if a.Parent == x && b.Parent == x && c.Parent == x && d.Parent == x {
		fmt.Println("parents of a, b, c and d should be the same as x")
	}

	// Output:
	// union of a and b produces new root as b of rank=1
	// union of x and c produces new root as x of rank=1
	// union of y and d produces new root as x of rank=1
	// parents of a, b, c and d should be the same as x
}
