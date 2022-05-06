package binary

import "golang.org/x/exp/constraints"

// Search for elem in slice a,
// if elem exists, returns index
// else returns -1
func Search[T constraints.Ordered](a []T, elem T) int {

	l, m, h := 0, 0, len(a)-1

	for l <= h {
		m = (l + h) / 2
		if a[m] == elem {
			return m
		} else if a[m] > elem {
			h = m - 1
		} else {
			l = m + 1
		}
	}

	return -1
}
