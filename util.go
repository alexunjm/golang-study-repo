package main

// type intSlice []int

func (s intSlice) elementsEqual(to intSlice) bool {
	if len(s) != len(to) {
		return false
	}
	for i, value := range s {
		if to[i] != value {
			return false
		}
	}
	return true
}
