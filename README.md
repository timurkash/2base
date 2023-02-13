```bash
goos: linux
goarch: amd64
pkg: vscode2/2base
cpu: Intel(R) Core(TM) i5-7500 CPU @ 3.40GHz
=== RUN   Benchmark2Base
Benchmark2Base
Benchmark2Base-4        1000000000               0.6612 ns/op          0 B/op          0 allocs/op
=== RUN   Benchmark2Base_2
Benchmark2Base_2
Benchmark2Base_2-4
1000000000               0.5179 ns/op          0 B/op          0 allocs/op
=== RUN   Benchmark2Base2
Benchmark2Base2
Benchmark2Base2-4       92384128                15.76 ns/op            0 B/op          0 allocs/op
=== RUN   Benchmark2Base3
Benchmark2Base3
Benchmark2Base3-4       93971832                16.16 ns/op            0 B/op          0 allocs/op
=== RUN   Benchmark2Base3_2
Benchmark2Base3_2
Benchmark2Base3_2-4     93747319                16.50 ns/op            0 B/op          0 allocs/op
=== RUN   Benchmark2Base4
Benchmark2Base4
Benchmark2Base4-4       53225875                21.97 ns/op            0 B/op          0 allocs/op
PASS
ok      vscode2/2base   7.082s
```