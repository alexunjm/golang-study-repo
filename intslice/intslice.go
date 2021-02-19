package intslice

/*

	public

*/

// IntSlice shortcut for []int
type IntSlice []int

// Sort elements
func (s IntSlice) Sort() IntSlice {

	return sort(s)
}

/*

	private

*/

func sort(sliceArg IntSlice) IntSlice {

	var (
		slicePointerValue int
		pointer           = 1
		pointerLeft       int
	)

	slice := make(IntSlice, len(sliceArg))
	copy(slice, sliceArg)

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

func movement(slice IntSlice, found int, position int) {

	if found == position {
		return
	}

	value := slice[found]

	for pointer := found; pointer > position; pointer-- {
		slice[pointer] = slice[pointer-1]
	}

	slice[position] = value
}

func reverse(slice IntSlice) IntSlice {

	s := make(IntSlice, len(slice))
	copy(s, slice)

	var opp int
	for pointer := len(s)/2 - 1; pointer >= 0; pointer-- {
		opp = len(s) - 1 - pointer
		s[pointer], s[opp] = s[opp], s[pointer]
	}

	return s
}
