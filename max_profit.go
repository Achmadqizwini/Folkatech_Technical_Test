package main

import (
	"fmt"
	"math"
	"sort"
)

func findMaxProfit(price []int, k int) int {
	var transaction = [][]int{}
	var resultArr = []int{}
	var resultSum = 0
	for i := 0; i < len(price); i++ {
		highestPrice := 0
		for sell := i + 1; sell < len(price); sell++ {

			if price[sell] < price[sell-1] {
				highestPrice = price[sell-1]
			} else if sell == len(price)-1 {
				highestPrice = price[sell]
			}

			if price[sell] < highestPrice {
				break
			}

		}
		transaction = append(transaction, []int{price[i], highestPrice})
	}
	var bestToSell = map[int]bool{}
	for idx := 0; idx < len(transaction)-1; idx++ {
		if bestToSell[transaction[idx][1]] {
			continue
		}
		bestToSell[transaction[idx][1]] = true

		resultArr = append(resultArr, transaction[idx][0], transaction[idx][1])
	}

	var sellingPrice = []int{}
	for i := 0; i < len(resultArr)-1; i++ {
		diff := math.Abs(float64(resultArr[i] - resultArr[i+1]))
		if (i+1)%2 == 0 {
			continue
		}
		sellingPrice = append(sellingPrice, int(diff))
	}

	sort.SliceStable(sellingPrice, func(i, j int) bool {
		return sellingPrice[i] > sellingPrice[j]
	})

	for numOfTx, v := range sellingPrice {
		if numOfTx+1 > k {
			break
		} else {
			resultSum += v
		}
	}

	fmt.Println(resultArr, "\nmax of transaction:", k)
	return resultSum
}

func main() {
	fmt.Println(findMaxProfit([]int{4, 11, 2, 20, 59, 80}, 2))
	fmt.Println(findMaxProfit([]int{4, 11, 2, 20, 59, 80, 4, 20}, 3))
	fmt.Println(findMaxProfit([]int{4, 11}, 2))
	fmt.Println(findMaxProfit([]int{11, 0, 1, 5}, 2))
}
