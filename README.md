# Make slices in Go compatison

Comparison between making Go slices with a given capacity or not, and then appending to it. When no capacity is given, several copies need to be made to the slice when the capacity is reached.

Benchmark results:

```txt
goos: linux
goarch: amd64
pkg: makeslices
cpu: Intel(R) Core(TM) i7-10750H CPU @ 2.60GHz
BenchmarkMakeSlicesNoCap/len_=_0-12             808363870                1.507 ns/op
BenchmarkMakeSlicesNoCap/len_=_1-12             15897050                78.26 ns/op
BenchmarkMakeSlicesNoCap/len_=_2-12              6779947               185.5 ns/op
BenchmarkMakeSlicesNoCap/len_=_5-12              1859090               639.3 ns/op
BenchmarkMakeSlicesNoCap/len_=_10-12             1000000              1211 ns/op
BenchmarkMakeSlicesNoCap/len_=_20-12              493910              2436 ns/op
BenchmarkMakeSlicesNoCap/len_=_50-12              251532              5145 ns/op
BenchmarkMakeSlicesNoCap/len_=_100-12             157482              7674 ns/op
BenchmarkMakeSlicesNoCap/len_=_200-12              68462             20654 ns/op
BenchmarkMakeSlicesNoCap/len_=_1000-12             18279             60638 ns/op
BenchmarkMakeSlicesNoCap/len_=_10000-12             1028           1347286 ns/op
BenchmarkMakeSlicesNoCap/len_=_100000-12              66          18775088 ns/op
BenchmarkMakeSlicesNoCap/len_=_1000000-12              1        1079651700 ns/op
BenchmarkMakeSlicesCap/len_=_0-12               157862685                9.437 ns/op
BenchmarkMakeSlicesCap/len_=_1-12               18113254                60.91 ns/op
BenchmarkMakeSlicesCap/len_=_2-12               14757152                90.88 ns/op
BenchmarkMakeSlicesCap/len_=_5-12                6043624               197.8 ns/op
BenchmarkMakeSlicesCap/len_=_10-12               3343796               368.7 ns/op
BenchmarkMakeSlicesCap/len_=_20-12               1951034               679.9 ns/op
BenchmarkMakeSlicesCap/len_=_50-12                860998              1712 ns/op
BenchmarkMakeSlicesCap/len_=_100-12               317551              3362 ns/op
BenchmarkMakeSlicesCap/len_=_200-12               239186              5913 ns/op
BenchmarkMakeSlicesCap/len_=_1000-12               38236             29257 ns/op
BenchmarkMakeSlicesCap/len_=_10000-12               4333            278286 ns/op
BenchmarkMakeSlicesCap/len_=_100000-12               376           3063873 ns/op
BenchmarkMakeSlicesCap/len_=_1000000-12               49          25205371 ns/op
```

When capacity is given, we see speed-ups of up to 10x.
