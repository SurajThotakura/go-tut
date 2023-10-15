# Testing in Go

Testing comes out of the box in Go, no need to depend on external libraries

To write a test we just write a function with the following rules,

- It needs to be in a file with a name like xxx_test.go
- The test function must start with the word Test
- The test function takes one argument only (t *testing.T)
- "testing.T" comes from a package called "testing" which should be included in this file
