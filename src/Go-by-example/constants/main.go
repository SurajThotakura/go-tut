package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {

	fmt.Println(s)

	const n = 20000000000

	fmt.Println(int64(math.Sin(n)))

}
