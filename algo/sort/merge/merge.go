package merge

// Sort1 divides by creating new slices
// - recursive
func Sort1(a []int) []int {
	if len(a) > 1 {
		n := len(a) / 2
		return Merge1(Sort1(a[:n]), Sort1(a[n:]))
	}
	return a
}

// Merge1 merges into a new slice
func Merge1(l, r []int) []int {
	tl := len(l) + len(r)
	i, li, ri := 0, 0, 0
	res := make([]int, 0, tl)

	for i = 0; i < tl; i++ {
		if li == len(l) {
			res = append(res, r[ri:]...)
			break
		} else if ri == len(r) {
			res = append(res, l[li:]...)
			break
		} else if l[li] <= r[ri] {
			res = append(res, l[li])
			li++
		} else {
			res = append(res, r[ri])
			ri++
		}
	}
	return res
}

var ls, rs []int

// Sort2 divides by passing indexes.
// - recursive
func Sort2(a []int) {
	r := len(a) / 2
	ls = make([]int, r+2, r+2)
	rs = make([]int, r+2, r+2)
	sort2(a, 0, len(a))
}

// sort2 ,sort `a` from `p` to `q`
// - recursive
func sort2(a []int, p, q int) {
	if q-p > 1 {
		r := (p + q) / 2
		sort2(a, p, r)
		sort2(a, r, q)
		Merge2(a, p, r, q)
	}
}

func Sort3(a []int) []int {
	r := len(a) / 2
	ls = make([]int, len(a), len(a))
	rs = make([]int, r+2, r+2)
	sort3(a, 0, len(a))
	return a
}

// for loop faster than recursion
//
func sort3(a []int, p, q int) {
	step, mid, end := 1, 0, 0
	for ; step < q; step += step {
		for ls := 0; ls < q; ls += 2 * step {
			mid = min(ls+step, q) // mid should be calculated with step
			// not through (ls+end)/2 as it considers `rs`
			// to be sorted, which may not be the case.
			end = min(ls+(2*step), q)
			Merge2(a, ls, mid, end)
		}
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// Merge2
func Merge2(a []int, l, m, n int) {
	ll, rl := m-l, n-m

	copyTo(l, m, a, 0, ls)
	copyTo(m, n, a, 0, rs)

	li, ri := 0, 0

	for ; l < n; l++ {
		if li == ll {
			copyTo(ri, rl, rs, l, a)
			break
		} else if ri == rl {
			copyTo(li, ll, ls, l, a)
			break
		}
		if ls[li] < rs[ri] {
			a[l] = ls[li]
			li++
		} else {
			a[l] = rs[ri]
			ri++
		}
	}
}

func copyTo(from, to int, src []int, init int, dst []int) {
	//	fmt.Fprintln(f, from, to, init)
	for ; from < to; from, init = from+1, init+1 {
		dst[init] = src[from]
	}
}
