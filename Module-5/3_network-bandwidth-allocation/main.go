package main

import "slices"

type DataStream struct {
	id        int
	bandwidth int
	priority  int
}

func BandwidthAllocation(dataStreams []DataStream, routers []int) (routerAllocations map[int][]DataStream, streamsAllocated int) {
	routersCopy := slices.Clone(routers) // so we dont modify original
	routerAllocations = make(map[int][]DataStream)
	streamsAllocated = 0

	slices.SortFunc(dataStreams, func(a, b DataStream) int {
		return b.priority - a.priority
	})

	for _, stream := range dataStreams {
		bestFit := -1

		// find best-fit router
		for i := 0; i < len(routersCopy); i++ {
			if stream.bandwidth > routersCopy[i] {
				continue // doesnt fit
			}

			if bestFit == -1 || routers[i] < routers[bestFit] {
				bestFit = i // better fit found
			}
		}

		// if possible, allocate stream to best router
		if bestFit != -1 {
			routerAllocations[bestFit] = append(routerAllocations[bestFit], stream)
			routersCopy[bestFit] -= stream.bandwidth
			streamsAllocated++
		}
	}

	return routerAllocations, streamsAllocated
}
