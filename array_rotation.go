package main

import (
	"fmt"
	"math"
)

func arrayRotation(ar []int, shift int) []int {
	n := int(math.Sqrt(float64(len(ar))))
	// fmt.Println(ar)
	start := 1
	for start <= shift {
		for i := start; i < n; i++ {
			j := shift - 1
			ar[i], ar[j] = ar[j], ar[i]

		}

		// fmt.Println("===============")
		for i := 2*n - 1; i < len(ar); i += n {
			j := shift - 1
			ar[i], ar[j] = ar[j], ar[i]
		}

		// fmt.Println("===============")
		for i := len(ar) - 2; i >= len(ar)-n; i-- {
			j := shift - 1
			ar[i], ar[j] = ar[j], ar[i]
		}

		// fmt.Println("===============")
		for i := len(ar) - 2*n; i >= 0; i -= n {
			j := shift - 1
			ar[i], ar[j] = ar[j], ar[i]
		}
		fmt.Println("ROTATE", start, "TIMES", ar)
		start++
	}
	return ar
}

func main() {
	fmt.Println(arrayRotation([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 2))
	// fmt.Println(arrayRotation([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 2))
}
