# Solutions for advent of code 2022

https://adventofcode.com/2022

## benchmarking go code

Add a function with a name starting with `Bench` to a `*_test.go` file. Something like this:

```go
func BenchmarkDay1(b *testing.B) {
	b.Run("label of benchmark", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			functionToBenchmark()
		}
	})
}
```

Then add the `-bench` flag to the `go test` command, optionally adding `-benchmem`

```bash
go test -bench=. -benchmem
```

## profiling go code

set up benchmarks as above then generate and view a cpu profile like so:

```bash
go test -bench=. -cpuprofile=before.prof
go tool pprof -http localhost:3000 before.prof
```
