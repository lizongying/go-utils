package utils

import (
	"testing"
)

// go test -bench='BenchmarkInSlice$' -count=10 .
// goos: darwin
// goarch: arm64
// pkg: github.com/lizongying/go-utils
// BenchmarkInSlice-10            5         207184892 ns/op
// BenchmarkInSlice-10            5         210529525 ns/op
// BenchmarkInSlice-10            5         204624450 ns/op
// BenchmarkInSlice-10            5         205362650 ns/op
// BenchmarkInSlice-10            5         203210925 ns/op
// BenchmarkInSlice-10            5         204164450 ns/op
// BenchmarkInSlice-10            5         201273425 ns/op
// BenchmarkInSlice-10            5         208560850 ns/op
// BenchmarkInSlice-10            5         207711300 ns/op
// BenchmarkInSlice-10            5         206910067 ns/op
// PASS
// ok      github.com/lizongying/go-utils  22.205s
func BenchmarkInSlice(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var a []int
		for i := 1; i < 100000000; i++ {
			a = append(a, i)
		}
		InSlice(100000000, a)
	}
}
