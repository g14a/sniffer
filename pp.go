package main

import (
	"fmt"
	"log"
	"time"

	"github.com/fatih/color"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
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
	// Let's see if the packet is an ethernet packet
	ethernetLayer := packet.Layer(layers.LayerTypeEthernet)
	if ethernetLayer != nil {
		ethernetPacket, _ := ethernetLayer.(*layers.Ethernet)

		side = color.New(color.FgHiMagenta)
		value = color.New(color.FgYellow)

		side.Print("DST: ")
		value.Print(ethernetPacket.DstMAC, "     | ")

		side.Print("SRC: ")
		value.Print(ethernetPacket.SrcMAC, "     |")

		// Ethernet type is typically IPv4 but could be ARP or other
		side.Print("Type: ")
		value.Print(ethernetPacket.EthernetType)

		fmt.Println('\n')
	}

	// Let's see if the packet is IP (even though the ether type told us)
	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	if ipLayer != nil {
		ip, _ := ipLayer.(*layers.IPv4)

		side = color.New(color.FgCyan)
		value = color.New(color.FgYellow)

		side.Print("IP Version: ")
		value.Print(ip.Version, " | ")
		side.Print("IHL: ")
		value.Print(ip.IHL, "   | ")
		side.Print("TOS: ")
		value.Print(ip.TOS, " | ")
		side.Print("Length: ")
		value.Print(ip.Length, " 	  |\n")
		side.Print("     ID: ")
		value.Print(ip.Id)
		value.Print("           | ")
		side.Print("Flags: ")
		value.Print(ip.Flags, "  | ")
		side.Print("OFF: ")
		value.Print(ip.FragOffset, " 	   | ")
		fmt.Println()
		side.Print("TTL: ")
		value.Print(ip.TTL, "       | ")
		side.Print("Pro: ")
		value.Print(ip.Protocol, "  | ")
		side.Print("	CHS: ")
		value.Print(ip.Checksum, "   	   | \n")
		side.Print("		Src: ")
		value.Print(ip.SrcIP, "	    	   | \n")
		side.Print("	        Dst: ")
		value.Print(ip.DstIP, " 		   |  ")

		fmt.Println("\n\n")
	}
}
