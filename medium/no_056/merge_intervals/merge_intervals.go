package merge_intervals

import "sort"

type Interval struct {
	Start int
	End   int
}

func merge(intervals []Interval) []Interval {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})
	result := []Interval{}
	for _, value := range intervals {
		if len(result) == 0 {
			result = append(result, value)
			continue
		}
		pre := result[len(result)-1]
		if pre.End >= value.Start {
			if result[len(result)-1].End < value.End {
				result[len(result)-1].End = value.End
			}
		} else {
			result = append(result, value)
		}
	}
	return result
}
