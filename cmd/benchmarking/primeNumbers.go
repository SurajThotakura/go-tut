package benchmarking

import (
	"strconv"
)

func checkPrimeNumber(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func printPrimesTill(num int) string {
	store := ""
	for i := 1; i < num; i++ {
		if checkPrimeNumber(i) {
			store += (strconv.Itoa(i) + ", ")
		}
	}
	return store[:len(store)-2]
}
