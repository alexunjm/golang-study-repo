package ex2

func mergeSort(s []int) []int {
	size := len(s)
	if size < 2 {
		return s
	}
	resolve(s, (size+1)/2, size)
	return s
}

func resolve(s []int, midIndex int, size int) []int {
	sizeL := len(s[:midIndex])
	if sizeL > 1 {
		resolve(s[:midIndex], (sizeL+1)/2, sizeL)
	}
	sizeR := len(s[midIndex:])
	if sizeR > 1 {
		resolve(s[midIndex:], (sizeR+1)/2, sizeR)
	}
	pl, pr, i := 0, 0, 0
	sorted := make([]int, sizeL+sizeR)
	for {

		if i >= len(sorted) {
			break
		}
		if pr >= sizeR {
			sorted[i] = s[:midIndex][pl]
			pl++
			i++
			continue
		}
		if pl >= sizeL {
			sorted[i] = s[midIndex:][pr]
			pr++
			i++
			continue
		}

		if s[:midIndex][pl] < s[midIndex:][pr] {
			sorted[i] = s[:midIndex][pl]
			pl++
		} else {
			sorted[i] = s[midIndex:][pr]
			pr++
		}
		i++
	}
	subSet := append(s[0:0], sorted...)
	return subSet
}
