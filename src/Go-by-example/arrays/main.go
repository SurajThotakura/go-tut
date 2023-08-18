package main

import "fmt"

func main() {
	var a = [5]any{nil, 0, "apple", nil, 44}
	fmt.Println(a)

	var b [5]int
	for i := 0; i < 5; i++ {
		b[i] = (i + 1) / 2
	}
	fmt.Println(b)

	var array [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			array[i][j] = i + j
		}
	}
	fmt.Println(array)
}
