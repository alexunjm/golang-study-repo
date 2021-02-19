package intslice

/*

	public

*/

// MyIntSlice shortcut for []int
type MyIntSlice []int

// Sort elements
func Sort(sliceArg MyIntSlice) MyIntSlice {

	var (
		slicePointerValue int
		pointer           = 1
		pointerChange     int
		pointerLeft       int
		pointerRight      int
	)

	slice := make(MyIntSlice, len(sliceArg))
	copy(slice, sliceArg)

	for pointer, slicePointerValue = range slice {

		pointerRight = pointer - 1

		if pointerRight < 0 {
			continue
		}
		pointerLeft = 0
		for {
			if pointerRight < 0 {
				pointerChange = 0
				break
			}
			if slice[pointerRight] > slicePointerValue {
				pointerRight--
				if slice[pointerLeft] > slicePointerValue {
					if pointerLeft <= 0 {
						pointerChange = 0
					}
					pointerChange = pointerLeft
					break
				}
				continue
			} else {
				pointerChange = pointerRight + 1
				break
			}
		}
		movement(slice, pointer, pointerChange)
	}

	return slice
}

/*

	private

*/

func movement(slice MyIntSlice, found int, position int) {

	if found == position {
		return
	}

	value := slice[found]

	for pointer := found; pointer > position; pointer-- {
		slice[pointer] = slice[pointer-1]
	}

	slice[position] = value
}

func reverse(slice MyIntSlice) MyIntSlice {

	s := make(MyIntSlice, len(slice))
	copy(s, slice)

	var opp int
	for pointer := len(s)/2 - 1; pointer >= 0; pointer-- {
		opp = len(s) - 1 - pointer
		s[pointer], s[opp] = s[opp], s[pointer]
	}

	return s
}
