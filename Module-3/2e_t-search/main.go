package main

import (
	"fmt"
	"log"
	"math/rand"
	"sort"
)

func main() {
	A := randomArray(1000, 1000)
	sort.Ints(A)

	toFind := A[22]
	result, err := tSearch(A, toFind, 0, len(A))
	if err != nil {
		log.Fatalf("fuck")
	}

	fmt.Printf("k: %d, found: %d", toFind, A[result])
}

func tSearch(A []int, k, p, r int) (int, error) {
	if p <= r {
		q1 := p + (r-p)/3
		q2 := r - (r-p)/3

		if A[q1] == k {
			return q1, nil
		}
		if A[q2] == k {
			return q2, nil
		}

		if k < A[q1] {
			return tSearch(A, k, p, q1-1)
		} else if k > A[q2] {
			return tSearch(A, k, q2+1, r)
		} else {
			return tSearch(A, k, q1+1, q2-1)
		}
	}
	return 0, fmt.Errorf("not found")
}

func randomArray(length, maxVal int) (outputArray []int) {
	for i := 0; i < length; i++ {
		outputArray = append(outputArray, rand.Intn(maxVal))
	}
	return
}
