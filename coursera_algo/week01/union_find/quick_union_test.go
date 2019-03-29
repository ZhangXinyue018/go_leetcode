package union_find

import (
	"fmt"
	"testing"
)

func TestQuickUnion_General(t *testing.T) {
	var uf UnionFind
	uf = BuildQuickUnion(10)
	type datapair struct {
		p int
		q int
	}
	dataInput := []datapair{
		{1, 2}, {3, 4}, {5, 6}, {0, 5}, {2, 8}, {1, 9},
	}

	for _, data := range dataInput{
		err := uf.Union(data.p, data.q)
		if err != nil{
			t.Errorf("failure detected doing uion for %v", data)
		}
	}
	fmt.Println(uf)
	if result, _ := uf.Connected(0, 6); !result {
		t.Errorf("wrong implementation")
	}
}