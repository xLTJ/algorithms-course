package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arrayToSort := randomArray(10, 10)
	fmt.Println("Original Array: ", arrayToSort)
	insertionSort(arrayToSort)
	fmt.Println("Sorted Array: ", arrayToSort)
}

func insertionSort(A []int) {
	n := len(A)
	for j := 1; j < n; j++ {
		key := A[j]
		i := j - 1
		for i >= 0 && A[i] > key {
			A[i+1] = A[i]
			i = i - 1
		}
		A[i+1] = key
	}
}

func randomArray(length, maxVal int) (outputArray []int) {
	for i := 0; i < length; i++ {
		outputArray = append(outputArray, rand.Intn(maxVal))
	}
	return
}
