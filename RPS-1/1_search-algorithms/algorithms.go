package main

import "fmt"

func LSearch(A []int, v int) (int, error) {
	n := len(A)
	for i := 0; i < n; i++ {
		if A[i] == v {
			return i, nil
		}
	}
	return -1, fmt.Errorf("target %d was not found in array", v)
}

func BSearch(A []int, v int) (int, error) {
	return BSearchHelper(A, 0, len(A)-1, v)
}

func BSearchHelper(A []int, l, r, v int) (int, error) {
	for l <= r {
		m := l + (r-l)/2
		if A[m] == v {
			return m, nil
		} else if A[m] < v {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return -1, fmt.Errorf("target %d was not found in array", v)
}

func FSearch(A []int, v int) (int, error) {
	return FSearchHelper(A, 0, len(A)-1, v)
}

func FSearchHelper(A []int, l, r, v int) (int, error) {
	if l > r {
		return -1, fmt.Errorf("target %d was not found in array", v)
	}

	m1 := l + (r-l)/4
	m2 := l + 2*(r-l)/4
	m3 := l + 3*(r-l)/4

	switch {
	case v == A[m1]:
		return m1, nil
	case v == A[m2]:
		return m2, nil
	case v == A[m3]:
		return m3, nil
	case v < A[m1]:
		return FSearchHelper(A, l, m1-1, v)
	case v < A[m2]:
		return FSearchHelper(A, m1+1, m2-1, v)
	case v < A[m3]:
		return FSearchHelper(A, m2+1, m3-1, v)
	default:
		return FSearchHelper(A, m3+1, r, v)
	}
}
