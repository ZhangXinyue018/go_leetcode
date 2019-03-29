package union_find

type UnionFind interface {
	// union points
	Union(p int, q int) error

	// check if p and q are connected
	Connected(p int, q int) (bool, error)

	// no of total components
	Count() int
}

