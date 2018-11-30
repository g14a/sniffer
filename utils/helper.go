package utils

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

var (
	side  *color.Color
	value *color.Color
)

func PPEthernetPacket(packet gopacket.Packet) {

	ethernetLayer := packet.Layer(layers.LayerTypeEthernet)

	if ethernetLayer != nil {
		ethernetPacket, _ := ethernetLayer.(*layers.Ethernet)

		side = color.New(color.FgHiMagenta)
		value = color.New(color.FgYellow)

		fmt.Println("\nEthernet Layer Found!")

		side.Print("DST: ")
		value.Print(ethernetPacket.DstMAC, "     | ")

		side.Print("	SRC: ")
		value.Print(ethernetPacket.SrcMAC, "     |")

		// Ethernet type is typically IPv4 but could be ARP or other
		side.Print("  Type: ")
		value.Print(ethernetPacket.EthernetType)

		fmt.Println()
	}
}

func PPTcpPacket(packet gopacket.Packet) {
	tcpLayer := packet.Layer(layers.LayerTypeTCP)

	if tcpLayer != nil {
		tcpPacket, _ := tcpLayer.(*layers.TCP)

		side = color.New(color.FgCyan)
		value = color.New(color.FgYellow)

		fmt.Println("TCP Layer Found!")

		side.Print("SrcPort: ")
		value.Print(tcpPacket.SrcPort)
		side.Print("	DstPort: ")
		value.Print(tcpPacket.DstPort)
		fmt.Println()
		side.Print("	SEQ: ")
		value.Print(tcpPacket.Seq)
		fmt.Println()
		side.Print("	ACK: ")
		value.Print(tcpPacket.Ack)
		fmt.Println()
		side.Print("HLen: ")
		value.Print(tcpPacket.DataOffset)

		flagMap := make(map[string]bool, 9)

		flagMap["NS"] = tcpPacket.NS
		flagMap["CWR"] = tcpPacket.CWR
		flagMap["ECE"] = tcpPacket.ECE
		flagMap["URG"] = tcpPacket.URG
		flagMap["ACK"] = tcpPacket.ACK
		flagMap["PSH"] = tcpPacket.PSH
		flagMap["RST"] = tcpPacket.RST
		flagMap["SYN"] = tcpPacket.SYN
		flagMap["FIN"] = tcpPacket.FIN

		PPTCPFlags(flagMap)

		side.Print(" WinSize : ")
		value.Print(tcpPacket.Window, "\n")

		side.Print("Checksum: ")
		value.Print(tcpPacket.Checksum)

		side.Print("       URG PTR: ")
		value.Print(tcpPacket.Urgent, "\n")

		side.Print("	Options: ")
		options := make([]layers.TCPOption, 0)

		options = tcpPacket.Options

		value.Print(len(options))
	}
}

func PPIPPacket(packet gopacket.Packet) {
	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	if ipLayer != nil {
		ipPacket, _ := ipLayer.(*layers.IPv4)

		side = color.New(color.FgCyan)
		value = color.New(color.FgYellow)

		fmt.Println("IP Layer Found!")

		side.Print("IP Version: ")
		value.Print(ipPacket.Version)
		side.Print("	IHL: ")
		value.Print(ipPacket.IHL)
		side.Print("	TOS: ")
		value.Print(ipPacket.TOS)
		side.Print("	Length: ")
		value.Print(ipPacket.Length)
		side.Print("\n     ID: ")
		value.Print(ipPacket.Id)
		side.Print("  Flags: ")
		value.Print(ipPacket.Flags, "  ")
		side.Print("  OFF: ")
		value.Print(ipPacket.FragOffset)
		fmt.Println()
		side.Print("TTL: ")
		value.Print(ipPacket.TTL, "        ")
		side.Print("Pro: ")
		value.Print(ipPacket.Protocol, "   ")
		side.Print("	CHS: ")
		value.Print(ipPacket.Checksum, "   	   	  \n")
		side.Print("	    Src: ")
		value.Print(ipPacket.SrcIP, "	    	   	  \n")
		side.Print("	    Dst: ")
		value.Print(ipPacket.DstIP, " 		   	   ")

		fmt.Println()
	}
}

func PPTCPFlags(m map[string]bool) {

	side = color.New(color.FgCyan)
	value = color.New(color.FgYellow)

	side.Print(" Flags: ")

	for k, v := range m {
		if v {
			value.Print(k, ",")
		}
	}
}
