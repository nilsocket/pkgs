package binary

import (
	"golang.org/x/exp/constraints"
)

// Search for elem in slice a,
// if elem exists, returns index
// else returns -1
func Search[T constraints.Ordered](a []T, elem T) int {
	if len(a) == 0 {
		return -1
	}

	mid := len(a) / 2
	curElem := a[mid]

	if curElem == elem {
		return mid
	} else if curElem < elem {
		if r := Search(a[mid+1:], elem); r == -1 {
			return -1
		} else {
			return mid + 1 + r
		}
	} else {
		return Search(a[:mid], elem)
	}
}
