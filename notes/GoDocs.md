# Go Docs

```godoc -http :8000``` (or any vacant port number)

## Example function

By writing an example function with the following rules in the "_test.go" file the documentation automatically gets updated with the example of the current package.

- The example function should be named as "ExampleExactFunctionName"
- The function should end with a comment that gives an expected output of the function ```//Output: {output of the function}```
- If the comment is removed the function will be compiled but it won't be executed in the test and won't show up in the doc
