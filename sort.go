package main

type intSlice []int

/*
func sort(sliceArg intSlice) intSlice {
	var (
		slicePointerValue int
		pointer           = 1
		pointerLeft       int
	)
	// println(&sliceArg)
	slice := make(intSlice, len(sliceArg))
	copy(slice, sliceArg)
	// println(&slice)

	for pointer, slicePointerValue = range slice {

		pointerLeft = pointer - 1

		if pointerLeft < 0 {
			continue
		}
		for {
			if pointerLeft < 0 {
				pointerLeft = 0
				break
			}
			if slice[pointerLeft] > slicePointerValue {
				pointerLeft--
				continue
			} else {
				pointerLeft++
				break
			}
		}
		movement(slice, pointer, pointerLeft)
	}
	return slice
}
*/
func movement(slice intSlice, found int, position int) {
	if found == position {
		return
	}
	// front := slice[:position]
	value := slice[found]
	// move1 := slice[position:found]
	// move2 := slice[found+1:]
	for pointer := found; pointer > position; pointer-- {
		slice[pointer] = slice[pointer-1]
	}
	slice[position] = value
	// fmt.Printf(strconv.Itoa(found) + "m" + strconv.Itoa(position) + " ")
	// fmt.Println(slice)
	// prettyslice.Show(strconv.Itoa(found)+"m"+strconv.Itoa(position), slice)
}

func reverse(slice intSlice) intSlice {
	s := make(intSlice, len(slice))
	copy(s, slice)

	var opp int
	for pointer := len(s)/2 - 1; pointer >= 0; pointer-- {
		opp = len(s) - 1 - pointer
		s[pointer], s[opp] = s[opp], s[pointer]
	}

	return s
}
