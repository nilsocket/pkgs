package random

import "math/rand"

// Ints returns a random slice of ints with given size
func Ints(size int) []int {
	res := make([]int, size, size)
	size--
	for ; size > -1; size-- {
		res[size] = rand.Int()
	}
	return res
}

// Intsn returns a random slice of ints < n ,with given size
func Intsn(size, n int) []int {
	res := make([]int, size, size)
	size--
	for ; size > -1; size-- {
		res[size] = rand.Intn(n)
	}
	return res
}
