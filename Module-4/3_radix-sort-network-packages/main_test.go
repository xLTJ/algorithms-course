package main

import (
	"slices"
	"testing"
)

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
			packages, _ := parseInput(tt.args.inputFile)
			originalPackages := slices.Clone(packages) // for comparison later

			radixSortPackages(packages)

			t.Log("Testing if packages got sorted correctly...")
			prevTimestamp := packages[0].time
			for i, currentPackage := range packages {
				if currentPackage.time < prevTimestamp {
					t.Errorf("package %d has timestamp smaller than previous, packages not correctly sorted :c", i)
				}
				prevTimestamp = currentPackage.time
			}

			t.Log("Packages were sorted correctly, testing if any packages were lost...")

			packageMap := make(map[string]int) // turn it into a map so comparing is faster
			for _, networkPackage := range packages {
				packageMap[networkPackage.name] = networkPackage.time
			}

			for _, originalPackage := range originalPackages {
				foundPackageTime, ok := packageMap[originalPackage.name]
				if !ok {
					t.Errorf("Missing package: %v", originalPackage)
				}

				if foundPackageTime != originalPackage.time {
					t.Errorf("Package %s does not match original packet. Got: %v, expected %v", originalPackage.name, foundPackageTime, originalPackage.time)
				}
			}
		})
	}
}
