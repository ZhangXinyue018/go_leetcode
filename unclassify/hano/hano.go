package hano

import "fmt"

func HanoSub(from string, to string, numbers int) {
	if numbers == 1 {
		fmt.Printf("Move %s to %s\n", from, to)
		return
	}
	theOther := GetTheOther(from, to)
	HanoSub(from, theOther, numbers-1)
	HanoSub(from, to, 1)
	HanoSub(theOther, to, numbers-1)

}

func GetTheOther(from string, to string) string {
	strList := []string{"A", "B", "C"}
	for _, str := range strList {
		if str != from && str != to {
			return str
		}
	}
	return ""
}
