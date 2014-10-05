package flowtrack

import (
	"fmt"
	"net"
)

var (
	topFlowsPackets = map[flowkey]uint64{}
	topFlowsBytes   = map[flowkey]uint64{}

	topSourcePackets = map[addrPortKey]uint64{}
	topSourceBytes   = map[addrPortKey]uint64{}

	topDestPackets = map[addrPortKey]uint64{}
	topDestBytes   = map[addrPortKey]uint64{}
)

func Process(source, destination net.IP, sourcePort, destPort, bytes int) {
	flowKey := generateFlowKey(source, destination, sourcePort, destPort)
	sourceKey := generateAddrPortKey(source, sourcePort)
	destKey := generateAddrPortKey(destination, destPort)

	topFlowsPackets[flowKey] += 1
	topFlowsBytes[flowKey] += uint64(bytes)

	topSourcePackets[sourceKey] += 1
	topSourceBytes[sourceKey] += uint64(bytes)

	topDestPackets[destKey] += 1
	topDestBytes[destKey] += uint64(bytes)
}

func PrintTopN(n int) {
	keys := sortFlowKeyMap(topFlowsPackets)
	for i, key := range keys {
		if i >= n {
			break
		}
		fmt.Printf("%s [%d packets]\n", key, topFlowsPackets[key])
	}

	fmt.Println()

	addrKeys := sortAddrPortKeySortableMap(topSourcePackets)
	for i, key := range addrKeys {
		if i >= n {
			break
		}
		fmt.Printf("%s [%d packets]\n", key, topSourcePackets[key])
	}

	fmt.Println()

	addrKeys = sortAddrPortKeySortableMap(topDestPackets)
	for i, key := range addrKeys {
		if i >= n {
			break
		}

		fmt.Printf("%s [%d packets]\n", key, topDestPackets[key])
	}

	fmt.Println("total flows:", len(topFlowsBytes))
}

func Reset() {
	topFlowsPackets = map[flowkey]uint64{}
	topFlowsBytes = map[flowkey]uint64{}

	topSourcePackets = map[addrPortKey]uint64{}
	topSourceBytes = map[addrPortKey]uint64{}

	topDestPackets = map[addrPortKey]uint64{}
	topDestBytes = map[addrPortKey]uint64{}
}
