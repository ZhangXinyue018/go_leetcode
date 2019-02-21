package climbling_stairs

func climbStairs(n int) int {
	return ClimbRecursive(n)
}

func ClimbRecursive(n int) int {
	if n == 1 || n == 0 {
		return 1
	} else if n == 2 {
		return 2
	}

	if n%2 == 0 {
		return ClimbRecursive(n/2)*ClimbRecursive(n/2) + ClimbRecursive(n/2-1)*ClimbRecursive(n/2-1)
	} else {
		return ClimbRecursive(n/2)*ClimbRecursive(n/2+1) + ClimbRecursive(n/2)*ClimbRecursive(n/2-1)
	}
}
