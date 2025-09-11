package main

import (
	"slices"
	"sort"
	"testing"
)

var testArrays [][]int

func init() {
	for i := 0; i < 10; i++ {
		testArrays = append(testArrays, randomArray(1000, 1000))
	}
}

func Benchmark_insertionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arrayClone := slices.Clone(testArrays[i%len(testArrays)])
		insertionSort(arrayClone)
	}
}

func Test_insertionSort(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "10 values",
			args: args{A: []int{7, 2, 8, 2, 8, 9, 63, -3, 0, 1}},
		},
		{
			name: "1000 random values",
			args: args{A: randomArray(1000, 1000)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			insertionSort(tt.args.A)
			if !sort.IntsAreSorted(tt.args.A) {
				t.Fatalf("Test %s failed, array not sorted", tt.name)
			}
		})
	}
}
