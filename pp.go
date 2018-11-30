package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fatih/color"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"github.com/nsd/utils"
)

var (
	device      string
	snapshotLen int32 = 1024
	promiscuous bool  = false
	err         error
	timeout     time.Duration = 30 * time.Second
	handle      *pcap.Handle
)

var (
	side  *color.Color
	value *color.Color
)

func main() {
	// Open device for capture packets in non-promiscous mode.

	device = os.Args[1]

	handle, err = pcap.OpenLive(device, snapshotLen, promiscuous, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	count := 0
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	for packet := range packetSource.Packets() {
		count++
		printPacketInfo(packet, count)
	}
}

func printPacketInfo(packet gopacket.Packet, count int) {

	color.Green("\n---------------------------------------------------------------")

	fmt.Println("PACKET ", count)

	utils.PPEthernetPacket(packet)
	fmt.Println()

	utils.PPIPPacket(packet)
	fmt.Println()

	utils.PPTcpPacket(packet)

	color.Green("\n---------------------------------------------------------------")
}
