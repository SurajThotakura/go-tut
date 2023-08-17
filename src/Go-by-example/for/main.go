package main

import "fmt"

func main() {
	// i := 1
	// for i <= 3 {
	// 	fmt.Println(i)
	// 	i++
	// }

	for j := 0; j <= 10; j++ {

		if j%2 == 0 {
			continue
		}
		fmt.Println(j)
	}

	a := 0

	for {
		fmt.Println("hey")
		if a > 10 {
			break
		}
		a++
	}
}
