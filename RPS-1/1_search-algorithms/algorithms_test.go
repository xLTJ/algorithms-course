package main

import (
	"testing"
)

// SearchFunc defines the common interface for all search algorithms
type SearchFunc func([]int, int) (int, error)

func TestSearchAlgorithms(t *testing.T) {
	tests := []struct {
		name      string
		array     []int
		target    int
		wantIndex int
		wantErr   bool
	}{
		{
			name:      "found at beginning",
			array:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			target:    1,
			wantIndex: 0,
			wantErr:   false,
		},
		{
			name:      "found at end",
			array:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			target:    9,
			wantIndex: 8,
			wantErr:   false,
		},
		{
			name:      "found in middle",
			array:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			target:    5,
			wantIndex: 4,
			wantErr:   false,
		},
		{
			name:      "not found - too small",
			array:     []int{1, 2, 3, 4, 5},
			target:    0,
			wantIndex: -1,
			wantErr:   true,
		},
		{
			name:      "not found - too large",
			array:     []int{1, 2, 3, 4, 5},
			target:    99,
			wantIndex: -1,
			wantErr:   true,
		},
		{
			name:      "not found - in gap",
			array:     []int{1, 3, 5, 7, 9},
			target:    6,
			wantIndex: -1,
			wantErr:   true,
		},
		{
			name:      "empty array",
			array:     []int{},
			target:    5,
			wantIndex: -1,
			wantErr:   true,
		},
		{
			name:      "single element found",
			array:     []int{42},
			target:    42,
			wantIndex: 0,
			wantErr:   false,
		},
		{
			name:      "single element not found",
			array:     []int{42},
			target:    99,
			wantIndex: -1,
			wantErr:   true,
		},
		{
			name:      "two elements - first",
			array:     []int{10, 20},
			target:    10,
			wantIndex: 0,
			wantErr:   false,
		},
		{
			name:      "two elements - second",
			array:     []int{10, 20},
			target:    20,
			wantIndex: 1,
			wantErr:   false,
		},
		{
			name:      "large sorted array",
			array:     makeRange(1, 1000),
			target:    742,
			wantIndex: 741,
			wantErr:   false,
		},
	}

	// map of algorithm name to function
	algorithms := map[string]SearchFunc{
		"LSearch": LSearch,
		"BSearch": BSearch,
		"FSearch": FSearch,
	}

	// Run all test cases against all algorithms
	for algoName, algoFunc := range algorithms {
		t.Run(algoName, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					gotIndex, err := algoFunc(tt.array, tt.target)

					if (err != nil) != tt.wantErr {
						t.Errorf("%s() error = %v, wantErr %v", algoName, err, tt.wantErr)
						return
					}

					if gotIndex != tt.wantIndex {
						t.Errorf("%s() = %v, want %v", algoName, gotIndex, tt.wantIndex)
					}
				})
			}
		})
	}
}

// Special test for LSearch with unsorted array (only LSearch can handle this)
func TestLSearchUnsorted(t *testing.T) {
	tests := []struct {
		name      string
		array     []int
		target    int
		wantIndex int
		wantErr   bool
	}{
		{
			name:      "unsorted array - found",
			array:     []int{5, 2, 8, 1, 9},
			target:    8,
			wantIndex: 2,
			wantErr:   false,
		},
		{
			name:      "unsorted array - not found",
			array:     []int{5, 2, 8, 1, 9},
			target:    99,
			wantIndex: -1,
			wantErr:   true,
		},
		{
			name:      "duplicate values - returns first",
			array:     []int{3, 7, 7, 7, 9},
			target:    7,
			wantIndex: 1,
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIndex, err := LSearch(tt.array, tt.target)

			if (err != nil) != tt.wantErr {
				t.Errorf("LSearch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if gotIndex != tt.wantIndex {
				t.Errorf("LSearch() = %v, want %v", gotIndex, tt.wantIndex)
			}
		})
	}
}

// Helper function to create a range of integers
func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

// Benchmark comparison
func BenchmarkSearchComparison(b *testing.B) {
	sizes := []int{100, 1000, 10000, 10000000}

	for _, size := range sizes {
		arr := makeRange(1, size)
		target := size // Worst case: last element

		b.Run("LSearch_"+string(rune(size)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				LSearch(arr, target)
			}
		})

		b.Run("BSearch_"+string(rune(size)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				BSearch(arr, target)
			}
		})

		b.Run("FSearch_"+string(rune(size)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				FSearch(arr, target)
			}
		})
	}
}
