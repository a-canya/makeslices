package makeslices_test

import (
	"fmt"
	"makeslices"
	"testing"
)

var slicesLens = []int{0, 1, 2, 5, 10, 20, 50, 100, 200, 1000, 10_000, 100_000, 1_000_000}

func BenchmarkMakeSlicesNoCap(b *testing.B) {
	for _, slicesLen := range slicesLens {
		b.Run(fmt.Sprintf("len = %d", slicesLen), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				makeslices.MakeSlicesNoCap(slicesLen)
			}
		})
	}
}

func BenchmarkMakeSlicesCap(b *testing.B) {
	for _, slicesLen := range slicesLens {
		b.Run(fmt.Sprintf("len = %d", slicesLen), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				makeslices.MakeSlicesCap(slicesLen)
			}
		})
	}
}
