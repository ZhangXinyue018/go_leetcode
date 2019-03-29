package union_find

import "fmt"

type QuickFind struct {
	mapping []int
}

func BuildQuickFind(n int) *QuickFind {
	mapping := make([]int, n)
	for i := 0; i < n; i++ {
		mapping[i] = i
	}
	return &QuickFind{mapping: mapping}
}

func (qf *QuickFind) Union(p int, q int) error {
	if p >= qf.Count() || q >= qf.Count() {
		return fmt.Errorf("cannot union values %d %d, because they are not in data set", p, q)
	}

	pValue := qf.mapping[p]
	qValue := qf.mapping[q]
	if pValue == qValue {
		return nil
	}
	for i := 0; i < qf.Count(); i++ {
		if qf.mapping[i] == pValue {
			qf.mapping[i] = qValue
		}
	}
	return nil
}

func (qf *QuickFind) Count() int {
	return len(qf.mapping)
}

func (qf *QuickFind) Connected(p int, q int) (bool, error) {
	if p >= qf.Count() || q >= qf.Count() {
		return false, fmt.Errorf("cannot union values %d %d, because they are not in data set", p, q)
	}
	pValue := qf.mapping[p]
	qValue := qf.mapping[q]
	return pValue == qValue, nil
}
