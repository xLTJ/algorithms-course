package main

import "testing"

func Test_radixSortPackages(t *testing.T) {
	type args struct {
		inputFile string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "literally the only test",
			args: args{inputFile: "network packets data.csv"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			packets, _ := parseInput(tt.args.inputFile)
			radixSortPackages(packets)

			prevTimestamp := packets[0].time
			for i, currentPackage := range packets {
				if currentPackage.time < prevTimestamp {
					t.Errorf("package %d has timestamp smaller than previous, packages not correctly sorted :c", i)
				}
			}
		})
	}
}
