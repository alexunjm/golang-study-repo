package main

import (
	"fmt"
)

func main() {
	// ages, first and last2 have the same backing arrays
	ages := []int{5, 8, 20, 16, 3, 7, 9, 17}

	// prettyslice.Show("ages", ages)
	fmt.Printf("\n %v\n", ages)
	// sort(ages)
}
