# Benchmarking

It is useful to

- Compare the performance of different implementations
- Determine the impact of a change in the codebase
- Optimize code for specific use cases, and validate the scalability of the code

---

## Setup

They are written very similar to tests with `*testing.B` & `b.N` to run your particular function b.N times which is dynamically determined by the framework.

```Go
func BenchmarkPrimesTill(b *testing.B) {
    for i := 0; i < b.N; i++ {
        printPrimesTill(1000)
    }
}
```

---

### Reference

- [Practical blog article](https://blog.logrocket.com/benchmarking-golang-improve-function-performance/)