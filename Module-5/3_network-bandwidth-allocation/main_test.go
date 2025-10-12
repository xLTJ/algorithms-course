package main

import (
	"slices"
	"testing"
)

// I have come to the conclusion that this test has overcomplicated wayyy too much, but i spend too much time so its staying lmao
// i didnt wanna implement logic for prioity testing (cus then we need to check if any non-allocated streams violate the prioity.
// and see if the stream that violates the prioity thing could even fit if the ones with lower priority were gone?
// like wtf no way, im just checking ts manually
func Test_bandwidthAllocation(t *testing.T) {
	type args struct {
		dataStreams []DataStream
		routers     []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "the only test there is",
			args: args{
				dataStreams: []DataStream{
					{id: 1, bandwidth: 100, priority: 3},
					{id: 2, bandwidth: 200, priority: 5},
					{id: 3, bandwidth: 150, priority: 2},
					{id: 4, bandwidth: 80, priority: 4},
				},
				routers: []int{250, 300},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			routerAllocations, streamsAllocated := BandwidthAllocation(tt.args.dataStreams, tt.args.routers)
			for router, dataStreams := range routerAllocations {
				totalBandwidth := 0
				seenDataStreams := make(map[int]bool)

				for _, dataStream := range dataStreams {
					// check if data streams didnt get fucked somehow
					if !slices.Contains(tt.args.dataStreams, dataStream) {
						t.Errorf("unexpeted data stream: %v", dataStream)
					}

					// check if its a duplicate
					_, ok := seenDataStreams[dataStream.id]
					if ok {
						t.Errorf("Duplicate stream: %v", dataStream)
					}
					seenDataStreams[dataStream.id] = true

					// add to total bandwidth
					totalBandwidth += dataStream.bandwidth
				}

				// check if bandwidth is under capacity
				if totalBandwidth > tt.args.routers[router] {
					t.Errorf("bandwidth %d exceeds capacity of router %d", totalBandwidth, router)
				}

				for i := streamsAllocated; i < len(tt.args.dataStreams); i++ {
					// check if any more streams could be assigned to router.
					if tt.args.dataStreams[i].bandwidth+totalBandwidth <= tt.args.routers[router] {
						t.Errorf("stream %d (bandwidth: %d, priority: %d) could fit in router %d (used: %d/%d)",
							tt.args.dataStreams[i].id,
							tt.args.dataStreams[i].bandwidth,
							tt.args.dataStreams[i].priority,
							router,
							totalBandwidth,
							tt.args.routers[router])
					}
				}
			}
		})
	}
}
