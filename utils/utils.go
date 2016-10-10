package utils

import "fmt"

/** Test "For" loops
  various case for loop usage
*/
func ForTest() {
	fmt.Printf("\n==========================\nTest For control\n")
	var sum int = 0
	var i int
	// i = 0
	for i = 0; i < 10; i++ {
		sum += i
	}
	// 45, 10
	fmt.Printf("\t Test sum: %d [i: %d]\n", sum, i)

	sum = 0
	// i = 10
	for ; i > 0; i-- {
		sum += i
	}
	// 55, 0
	fmt.Printf("\t Test sum: %d [i: %d]\n", sum, i)

	sum = 0
	// i = 0
	for ; ; i = i + 1 {
		if i >= 10 {
			break
		}
		sum += i
	}
	// 45, 10
	fmt.Printf("\t Test sum: %d [i: %d]\n", sum, i)

	sum = 0
	// i = 10
	for i > 0 {
		i--
		sum += i
	}
	// 45, 0
	fmt.Printf("\t Test sum: %d [i: %d]\n", sum, i)

	sum = 0
	// i = 0
	for i < 10 {
		sum += i
		i++
	}
	// 45, 10
	fmt.Printf("\t Test sum: %d [i: %d]\n", sum, i)

	sum = 0
	// i = 10
	for {
		i--
		if i <= 0 {
			break
		}
		sum += i
	}
	// 45, 0
	fmt.Printf("\t Test sum: %d [i: %d]\n", sum, i)

	sum = 0
	// i = 0
	// Attention - i not changed! after loop
	for i := 0; i < 10; i++ {
		sum += i
	}
	// 45, 0
	fmt.Printf("\t Test sum without index changing: %d [i: %d]\n", sum, i)

	var arr1 [3]int
	arr1[0] = 20
	println("Array range:")
	for key, value := range arr1 {
		fmt.Printf("\t [%d] = %d\n", key, value)
	}

	// Reverse array
	for i, j := 0, len(arr1)-1; i < j; i, j = i+1, j-1 {
		arr1[i], arr1[j] = arr1[j], arr1[i]
	}
	fmt.Printf("Reverse array: %v \n", arr1)
}
