package main

import (
	"fmt"
	"log"
	"net"

	"github.com/vmorsell/udp-broadcast/common"
)

func main() {
	pc, err := net.ListenPacket(common.Network, fmt.Sprintf(":%d", common.Port))
	if err != nil {
		log.Fatalf("listen packet: %v", err)
	}
	defer pc.Close()

	buf := make([]byte, common.BufferSize)
	for {
		n, addr, err := pc.ReadFrom(buf)
		if err != nil {
			log.Fatalf("read from: %v", err)
		}
		fmt.Printf("%s: %s\n", addr, buf[:n])
	}
}
