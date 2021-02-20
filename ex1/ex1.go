package ex1

func test() int {
	for i := 0; i < 999; i++ {
		// 8, 6
		if result := ex1(i, 0); result > 8 {
			println(i, result)
			return i
		}
	}
	return -1
}

func ex1(n int, i int) int {
	if i == 0 {
		i = (n / 2) + (n % 2)
	}
	var sumR int
	for j := i + 1; j <= n; j++ {
		sumR += j
	}
	var sumL int
	for j := i - 1; j >= 0; j-- {
		sumL += j
	}
	if sumL == sumR {
		return i
	}
	i2 := i + ((n - i) / 2)
	if i == i2 {
		return -1
	}
	if sumL < sumR {
		return ex1(n, i2)
	}
	i2 = (n + i) / 2
	return ex1(n, i2)

}
