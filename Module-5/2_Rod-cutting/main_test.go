package main

import (
	"reflect"
	"testing"
)

func Test_cutRod(t *testing.T) {
	type args struct {
		prices   []int
		length   int
		cutPrice int
	}
	tests := []struct {
		name  string
		args  args
		want  []int // revenueTable
		want1 []int // firstCutTable
	}{
		{
			name: "High cut price (c=10, n=4) - no cuts optimal",
			args: args{
				prices:   []int{0, 1, 5, 8, 9, 10, 17, 20, 24, 30},
				length:   4,
				cutPrice: 10,
			},
			want:  []int{0, 1, 5, 8, 9},
			want1: []int{0, 1, 2, 3, 4},
		},
		{
			name: "Zero cut price (c=0, n=4) - original problem",
			args: args{
				prices:   []int{0, 1, 5, 8, 9, 10, 17, 20, 24, 30},
				length:   4,
				cutPrice: 0,
			},
			want:  []int{0, 1, 5, 8, 10},
			want1: []int{0, 1, 2, 3, 2},
		},
		{
			name: "Small length (n=1) - base case",
			args: args{
				prices:   []int{0, 1, 5, 8, 9},
				length:   1,
				cutPrice: 5,
			},
			want:  []int{0, 1},
			want1: []int{0, 1},
		},
		{
			name: "Small length (n=2, c=1) - cut not worth it",
			args: args{
				prices:   []int{0, 1, 5, 8, 9},
				length:   2,
				cutPrice: 1,
			},
			want:  []int{0, 1, 5},
			want1: []int{0, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := cutRod(tt.args.prices, tt.args.length, tt.args.cutPrice)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cutRod() revenueTable = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("cutRod() firstCutTable = %v, want %v", got1, tt.want1)
			}
		})
	}
}
