package sqrt

func mySqrt(x int) int {
	return Hit(0, x/2+1, x)
}

func Hit(startNo int, endNo int, target int) int {
	if endNo-startNo <= 1 {
		if target < startNo*startNo {
			return startNo - 1
		} else if target < endNo*endNo {
			return endNo - 1
		} else {
			return endNo
		}
	}

	middleNo := (startNo + endNo) / 2
	middleResult := middleNo * middleNo
	if target == middleResult {
		return middleNo
	} else if target > middleResult {
		return Hit(middleNo+1, endNo, target)
	} else {
		return Hit(startNo, middleNo-1, target)
	}
}
