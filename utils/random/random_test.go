package random

import "testing"

var intsRes []int

func BenchmarkIntsHundred(b *testing.B)  { benchmarkIntsI(100, b) }
func BenchmarkIntsThousand(b *testing.B) { benchmarkIntsI(1000, b) }
func BenchmarkIntsLakh(b *testing.B)     { benchmarkIntsI(100000, b) }
func BenchmarkIntsMillion(b *testing.B)  { benchmarkIntsI(1000000, b) }
func BenchmarkIntsBillion(b *testing.B)  { benchmarkIntsI(1000000000, b) }

func benchmarkIntsI(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		intsRes = Ints(i)
	}
}

func BenchmarkIntsnHundred(b *testing.B)  { benchmarkIntsnI(100, b) }
func BenchmarkIntsnThousand(b *testing.B) { benchmarkIntsnI(1000, b) }
func BenchmarkIntsnLakh(b *testing.B)     { benchmarkIntsnI(100000, b) }
func BenchmarkIntsnMillion(b *testing.B)  { benchmarkIntsnI(1000000, b) }
func BenchmarkIntsnBillion(b *testing.B)  { benchmarkIntsnI(1000000000, b) }

func benchmarkIntsnI(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		intsRes = Intsn(i, 100000)
	}
}
