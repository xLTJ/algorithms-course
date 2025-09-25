package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const inputFilePath = "network packets data.csv"

type packet struct {
	name string
	time int
}

func main() {
	packets, err := parseInput(inputFilePath)
	if err != nil {
		log.Panic(fmt.Errorf("error parsing file: %v", err))
	}

	radixSortPackages(packets)
	fmt.Println(packets)
}

// radixSortPackages does radixsort stuff, iykyk
// Assumes that all timestamps are 6 digits in total. No, im not gonna make that dynamic
func radixSortPackages(packetSlice []packet) {
	for place := 1; place <= 100000; place *= 10 {
		countingSortPackages(packetSlice, place)
	}
}

// countingSortPackages sorts the packetSlice based on its timestamp, based on a specific place.
// place is basically like 1st digit = 1, 2nd digit = 10, 3rd digit = 100, and so on
func countingSortPackages(packetSlice []packet, place int) {
	n := len(packetSlice)
	output := make([]packet, n)
	auxiliary := make([]int, 10) // just 0 - 9 cus its digits

	// for each packet, get the target digit and add to the auxiliary to count occurrences.
	for i := range packetSlice {
		digit := (packetSlice[i].time / place) % 10
		auxiliary[digit]++
	}

	// make auxiliary count cumulative
	for i := 1; i < len(auxiliary); i++ {
		auxiliary[i] += auxiliary[i-1]
	}

	// now to place the elements
	for i := n - 1; i >= 0; i-- {
		digit := (packetSlice[i].time / place) % 10
		output[auxiliary[digit]-1] = packetSlice[i] // place it in the output based on digit. -1 cus 0-based index
		auxiliary[digit]--                          // to handle duplicates
	}

	copy(packetSlice, output) // copy back into packetSlice cus we love pure functions
}

func parseInput(filePath string) ([]packet, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func() { _ = file.Close() }()

	var parsedInput []packet
	scanner := bufio.NewScanner(file)
	scanner.Scan() // bye bye first line

	// actual scanning loop
	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, ",")

		// skips any funky lines (aka the last line)
		if len(splitLine) != 2 {
			continue
		}

		// extract fields and populate struct with them
		timestampString := strings.ReplaceAll(splitLine[1], ":", "")
		timestamp, err := strconv.Atoi(timestampString)
		if err != nil {
			return nil, fmt.Errorf("error parsing timestamp %s: %v", timestampString, err)
		}

		newPacket := packet{
			name: splitLine[0],
			time: timestamp,
		}

		// add to slice
		parsedInput = append(parsedInput, newPacket)
	}

	return parsedInput, nil
}
