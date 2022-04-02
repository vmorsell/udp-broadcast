package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/vmorsell/udp-broadcast/common"
)

func main() {
	t0 := time.Now()

	conn, err := net.ListenPacket(common.Network, ":0")
	if err != nil {
		log.Fatalf("listen packet: %v", err)
	}
	defer conn.Close()

	dst, err := net.ResolveUDPAddr(common.Network, fmt.Sprintf("%s:%d", common.BroadcastAddr, common.Port))
	if err != nil {
		log.Fatalf("resolve udp addr: %v", err)
	}

	for {
		msg := fmt.Sprint(time.Since(t0).Seconds())
		_, err = conn.WriteTo([]byte(msg), dst)
		if err != nil {
			log.Fatalf("write to: %v", err)
		}
		fmt.Printf("sent: %s\n", msg)
		time.Sleep(time.Second)
	}
}
