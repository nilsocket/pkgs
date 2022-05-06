package binary

import "golang.org/x/exp/constraints"

// Search for elem in slice a,
// if elem exists, returns index
// else returns -1
func Search[T constraints.Ordered](a []T, elem T) int {

	l, m, h := 0, 0, len(a)

	for l < h {
		m = int(uint(l+h) >> 1)
		if a[m] >= elem {
			h = m
		} else {
			l = m + 1
		}
	}

	return l
}
