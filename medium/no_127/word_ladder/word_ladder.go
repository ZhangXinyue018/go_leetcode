package word_ladder

import (
	"math"
)

// todo: to be optimized
func ladderLength(beginWord string, endWord string, wordList []string) int {
	result := DetailedLadderLength(beginWord, endWord, wordList)
	if result == math.MaxInt32 {
		return 0
	} else {
		return result
	}
}

func DetailedLadderLength(beginWord string, endWord string, wordList []string) int {
	if beginWord == endWord {
		return 1
	}
	min := math.MaxInt32
	for index, tempWord := range wordList {
		if DifferInOne(beginWord, tempWord) {
			newList := append(append([]string{}, wordList[:index]...), wordList[index+1:]...)
			tempLength := 1 + DetailedLadderLength(tempWord, endWord, newList)
			if tempLength < min {
				min = tempLength
			}
		}
	}
	return min
}

func DifferInOne(wordA string, wordB string) bool {
	if wordA == wordB {
		return false
	}
	count := 0
	for i := 0; i < len(wordA); i++ {
		if wordA[i] != wordB[i] {
			count++
		}
		if count > 1 {
			return false
		}
	}
	return count == 1
}

func ladderLengthBetter(beginWord string, endWord string, wordList []string) int {
	distanceMap := make([]int, len(wordList))
	currentCount := 2
	checkList := []string{beginWord}
	for len(checkList) != 0 {
		newCheckList := []string{}
		for _, checkTemp := range checkList {
			for index, secondCheckTemp := range wordList {
				if distanceMap[index] != 0 {
					continue
				} else if checkTemp == secondCheckTemp {
					distanceMap[index] = currentCount - 1
				} else if DifferInOne(checkTemp, secondCheckTemp) {
					if secondCheckTemp == endWord {
						return currentCount
					}
					distanceMap[index] = currentCount
					newCheckList = append(newCheckList, secondCheckTemp)
				}
			}
		}
		currentCount++
		checkList = newCheckList
	}
	return 0
}

func ladderLengthBetter2(beginWord string, endWord string, wordList []string) int {
	currentCount := 2
	checkList := []string{beginWord}
	checkMap := map[string]bool{}
	for _, word := range wordList {
		checkMap[word] = true
	}
	for len(checkList) != 0 {
		newCheckList := []string{}
		for _, checkTemp := range checkList {
			for secondCheckTemp := range checkMap {
				if DifferInOne(checkTemp, secondCheckTemp) {
					if secondCheckTemp == endWord {
						return currentCount
					}
					delete(checkMap, secondCheckTemp)
					newCheckList = append(newCheckList, secondCheckTemp)
				}
			}
		}
		currentCount++
		checkList = newCheckList
	}
	return 0
}
