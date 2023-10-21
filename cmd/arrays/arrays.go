package arrays

// func ArraySum(inputArray [5]int) int {
// 	sum := 0
// 	for i := 0; i < len(inputArray); i++ {
// 		sum += inputArray[i]
// 	}
// 	return sum
// }

// arrays are of fixed size after initiation hence their type takes in a constant array size

func ArraySum(inputArray [5]int) int {
	sum := 0
	for _, number := range inputArray {
		sum += number
	}
	return sum
}

//  Slices do not encode the size of the collection unlike arrays

func SliceSum(inputSlice []int) int {
	sum := 0
	for _, number := range inputSlice {
		sum += number
	}
	return sum
}

func SliceOfSums(inputSlice ...[]int) []int {
	inputLength := len(inputSlice)
	sumSlice := make([]int, inputLength)

	for i, subSlice := range inputSlice {
		sumSlice[i] = SliceSum(subSlice)
	}
	return sumSlice
}

// func main() {
// 	SliceOfSums([]int{1, 2, 10}, []int{43, 65, 10}, []int{1, 10})
// }
