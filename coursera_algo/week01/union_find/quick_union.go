package union_find

import "fmt"

type QuickUnion struct {
	mapping []int
	weight  []int
}

func BuildQuickUnion(n int) *QuickUnion {
	mapping := make([]int, n)
	weight := make([]int, n)
	for i := 0; i < n; i++ {
		mapping[i] = i
	}
	return &QuickUnion{mapping: mapping, weight: weight}
}

func (qu *QuickUnion) Connected(p int, q int) (bool, error) {
	if p >= qu.Count() || q >= qu.Count() {
		return false, fmt.Errorf("cannot union values %d %d, because they are not in data set", p, q)
	}
	pRoot := qu.mapping[p]
	qRoot := qu.mapping[q]
	pRoot = qu.root(p)
	qRoot = qu.root(q)

	return pRoot == qRoot, nil
}

func (qu *QuickUnion) Count() int {
	return len(qu.mapping)
}

func (qu *QuickUnion) root(p int) int {
	for p != qu.mapping[p] {
		qu.mapping[p] = qu.mapping[qu.mapping[p]]
		p = qu.mapping[p]
	}
	return p
}

func (qu *QuickUnion) Union(p int, q int) error {
	if p >= qu.Count() || q >= qu.Count() {
		return fmt.Errorf("cannot union values %d %d, because they are not in data set", p, q)
	}
	pRoot := qu.root(p)
	qRoot := qu.root(q)

	if pRoot != qRoot {
		if qu.weight[pRoot] > qu.weight[qRoot] {
			qu.mapping[qRoot] = pRoot
			qu.weight[pRoot] = qu.weight[pRoot] + 1
		} else {
			qu.mapping[pRoot] = qRoot
			qu.weight[qRoot] = qu.weight[qRoot] + 1
		}
	}
	return nil
}
