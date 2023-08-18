package main

import "fmt"

func main() {
	var s []any
	fmt.Println(s)

	var a = [4]interface{}{"apple", "ball", "cat", "dog"}

	type ArrayData struct {
		Number  int
		Text    string
		Checked bool
	}

	DataArray := [1]ArrayData{{1, "apple", false}}

	fmt.Println(DataArray)

	n := a[1:]

	fmt.Println(append(n, 4, 5))
}
