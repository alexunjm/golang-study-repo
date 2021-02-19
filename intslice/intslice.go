package intslice

/*

	public

*/

// MyIntSlice shortcut for []int
type MyIntSlice []int

// Sort elements
func Sort(sliceArg MyIntSlice) MyIntSlice {

	size := len(sliceArg)

	// startIndex := makeMoves(slice, size, size*2)

	// return slice[startIndex : startIndex+size]

	// startIndex :=
	makeMoves(sliceArg, 0, size)

	return sliceArg
}

func makeMoves(slice MyIntSlice, startIndex int, endIndex int) int {

	var (
		slicePointerValue int
		pointer           = 1
		pointerChange     int
		pointerLeft       int
		pointerRight      int
	)

	tmpSlice := slice[startIndex:endIndex]
	for pointer, slicePointerValue = range tmpSlice {

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
			if tmpSlice[pointerRight] > slicePointerValue {
				pointerRight--
				if tmpSlice[pointerLeft] > slicePointerValue {
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
		movement(tmpSlice, pointer, pointerChange)
	}

	return startIndex
}

/*

	private

*/

func movement(slice MyIntSlice, found int, position int) *MyIntSlice {

	if found == position {
		return &slice
	}

	value := slice[found]
	/*
		for pointer := found; pointer > position; pointer-- {
			slice[pointer] = slice[pointer-1]
		}
	*/
	result := append(slice[position+1:position+1], slice[position:found]...)
	slice[position] = value
	return &result
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
