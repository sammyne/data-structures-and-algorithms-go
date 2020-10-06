package dijkstra

import (
	"reflect"
	"testing"
)

func TestDijkstra(t *testing.T) {
	testVector := []struct {
		edges  []WeightedEdge
		src    int
		dst    int
		expect []WeightedEdge
	}{
		{
			[]WeightedEdge{
				{4, 5, 0.35},
				{5, 4, 0.35},
				{4, 7, 0.37},
				{5, 7, 0.28},
				{7, 5, 0.28},
				{5, 1, 0.32},
				{0, 4, 0.38},
				{0, 2, 0.26},
				{7, 3, 0.39},
				{1, 3, 0.29},
				{2, 7, 0.34},
				{6, 2, 0.40},
				{3, 6, 0.52},
				{6, 0, 0.58},
				{6, 4, 0.93},
			},
			0,
			6,
			[]WeightedEdge{
				{0, 2, 0.26},
				{2, 7, 0.34},
				{7, 3, 0.39},
				{3, 6, 0.52},
			},
		},
	}

	for i, c := range testVector {
		got, err := Dijkstra(c.edges, c.src, c.dst)
		if err != nil {
			t.Fatalf("#%d unexpected error: %v", i, err)
		}
		if !reflect.DeepEqual(c.expect, got) {
			t.Fatalf("#%d failed: expect %#v, got %#v", i, c.expect, got)
		}
	}
}
