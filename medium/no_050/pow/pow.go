package pow

func myPow(x float64, n int) float64 {
	if n <0{
		x = 1/x
		n = -n
	}
	if n ==1{
		return x
	}else if n == 0{
		return 1
	}
	if n%2 == 0{
		temp := myPow(x, n>>1)
		return temp * temp
	}else{
		temp := myPow(x, (n-1)>>1)
		return x*temp*temp
	}
}
