package group_anagram

import "sort"

func groupAnagrams(strs []string) [][]string {
	resultMap := map[string][]string{}
	for _, str := range strs{
		sortStr := GetSort(str)
		if _, exist := resultMap[sortStr]; exist{
			resultMap[sortStr] = append(resultMap[sortStr], str)
		}else{
			resultMap[sortStr] = []string{str}
		}
	}
	result := [][]string{}
	for _, value := range resultMap{
		result = append(result, value)
	}
	return result
}

func GetSort(str string)string{
	strNew := append([]uint8{}, str...)
	sort.Slice(strNew, func(i, j int) bool {
		return strNew[i] <= strNew[j]
	})
	return string(strNew)
}
