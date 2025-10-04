package main

import "fmt"

func main() {
	prices := []int{0, 1, 5, 8, 9, 10, 17, 20, 24, 30}
	printCutRodSolution(prices, 4, 10)
}

func cutRod(prices []int, length, cutPrice int) ([]int, []int) {
	revenueTable := make([]int, length+1)
	firstCutTable := make([]int, length+1)

	for j := 1; j <= length; j++ { // for increasing rod length j
		currentMaxRevenue := prices[j] // initial best solution is no cut made
		firstCutTable[j] = j

		// then see if making a cut is potentially better
		for i := 1; i < j; i++ {
			if currentMaxRevenue < prices[i]+revenueTable[j-i]-cutPrice {
				currentMaxRevenue = prices[i] + revenueTable[j-i] - cutPrice
				firstCutTable[j] = i // best cut location so far for length j
			}
		}

		revenueTable[j] = currentMaxRevenue
	}
	return revenueTable, firstCutTable
}

func printCutRodSolution(prices []int, length, cutPrice int) {
	revenueTable, firstCutTable := cutRod(prices, length, cutPrice)
	fmt.Printf("Total Revenue: %d\n", revenueTable[length])
	fmt.Println("Solution:")
	for length > 0 {
		fmt.Println(firstCutTable[length])
		length -= firstCutTable[length]
	}
}
