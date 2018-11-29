package main

import (
	"log"
	"time"

	"github.com/fatih/color"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"github.com/net/utils"
)

var (
	device      string = "wlp2s0"
	snapshotLen int32  = 1024
	promiscuous bool   = false
	err         error
	timeout     time.Duration = 30 * time.Second
	handle      *pcap.Handle
)

var (
	side  *color.Color
	value *color.Color
)

func main() {
	// Open device
	handle, err = pcap.OpenLive(device, snapshotLen, promiscuous, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		printPacketInfo(packet)
	}
}

func printPacketInfo(packet gopacket.Packet) {

	// utils.PPEthernetPacket(packet)

	// utils.PPIPPacket(packet)

	utils.PPTcpPacket(packet)

}
