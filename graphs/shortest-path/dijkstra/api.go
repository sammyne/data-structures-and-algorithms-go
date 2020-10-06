package dijkstra

import (
	"container/heap"
	"errors"
	"math"
)

type DistanceTo struct {
	Distance float64
	To       int
}

type PQ []DistanceTo

// Len is the number of elements in the collection.
func (pq PQ) Len() int {
	return len(pq)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (pq PQ) Less(i int, j int) bool {
	return pq[i].Distance < pq[j].Distance
}

// Swap swaps the elements with indexes i and j.
func (pq PQ) Swap(i int, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQ) Push(x interface{}) {
	*pq = append(*pq, x.(DistanceTo))
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	ell := len(old)
	*pq = old[:ell-1]

	return old[ell-1]
}

type WeightedEdge struct {
	From   int
	To     int
	Weight float64
}

func Dijkstra(edges []WeightedEdge, src, sink int) ([]WeightedEdge, error) {
	vertices := make(map[int]struct{})
	for _, v := range edges {
		vertices[v.From], vertices[v.To] = struct{}{}, struct{}{}
	}
	nVertex := len(vertices)

	G := make([][]WeightedEdge, nVertex)
	for _, v := range edges {
		G[v.From] = append(G[v.From], v)
	}

	distancesTo := make([]float64, nVertex)
	for i := range distancesTo {
		distancesTo[i] = math.MaxFloat64
	}
	distancesTo[src] = 0

	edgesTo := make([]WeightedEdge, nVertex)

	pq := PQ{DistanceTo{Distance: 0, To: src}}
	done := make([]bool, nVertex)
	for pq.Len() > 0 {
		x := heap.Pop(&pq).(DistanceTo)
		done[x.To] = true

		for _, v := range G[x.To] {
			if done[v.To] {
				continue
			}

			if d := distancesTo[v.From] + v.Weight; d < distancesTo[v.To] {
				distancesTo[v.To] = d
				heap.Push(&pq, DistanceTo{Distance: d, To: v.To})
				edgesTo[v.To] = v
			}
		}
	}

	if distancesTo[sink] == math.MaxFloat64 {
		return nil, errors.New("unreachable")
	}

	var out []WeightedEdge
	for t := sink; t != src; {
		out, t = append(out, edgesTo[t]), edgesTo[t].From
	}

	for i, j := 0, len(out)-1; i < j; i, j = i+1, j-1 {
		out[i], out[j] = out[j], out[i]
	}

	return out, nil
}
