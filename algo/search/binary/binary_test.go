package binary_test

import (
	"sort"
	"testing"

	binaryv0 "github.com/nilsocket/pkgs/algo/search/binary/v0"
	binaryv1 "github.com/nilsocket/pkgs/algo/search/binary/v1"
	binaryv2 "github.com/nilsocket/pkgs/algo/search/binary/v2"
	"github.com/nilsocket/pkgs/algo/sort/merge"
	"github.com/nilsocket/pkgs/utils/random"
)

const (
	maxNum       = 1e6
	maxLen       = 1e6
	maxSearchLen = 1e3
)

var (
	randInp         = merge.Sort3(random.Intsn(maxLen, maxNum))
	randSearchElems = random.Intsn(maxSearchLen, maxNum)
)

func BenchmarkSearchv0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for i := range randSearchElems {
			binaryv0.Search(randInp, randSearchElems[i])
		}
	}
}

func BenchmarkSearchv1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for i := range randSearchElems {
			binaryv1.Search(randInp, randSearchElems[i])
		}
	}
}

func BenchmarkSearchv2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for i := range randSearchElems {
			binaryv2.Search(randInp, randSearchElems[i])
		}
	}
}

func BenchmarkSearchStdlib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for i := range randSearchElems {
			sort.SearchInts(randInp, randSearchElems[i])
		}
	}
}
