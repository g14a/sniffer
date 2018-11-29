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

		side.Print("DST: ")
		value.Print(ethernetPacket.DstMAC, "     | ")

		side.Print("	SRC: ")
		value.Print(ethernetPacket.SrcMAC, "     	 |")

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

		fmt.Println("\n\n")

		side.Print("SrcPort: ")
		value.Print(tcpPacket.SrcPort)
		value.Print(" 		|")
		side.Print("	DstPort: ")
		value.Print(tcpPacket.DstPort)
		value.Print("    	|\n")
		side.Print("SEQ: ")
		value.Print(tcpPacket.Seq)
		value.Print(" 	|")
		side.Print("	ACK: ")
		value.Print(tcpPacket.Ack)
		value.Print(" 		|")
		side.Print("HLen: ")
		value.Print(tcpPacket.DataOffset)
		side.Print("RES: ")
		value.Print(tcpPacket.)
	}
}

func PPIPPacket(packet gopacket.Packet) {
	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	if ipLayer != nil {
		ip, _ := ipLayer.(*layers.IPv4)

		side = color.New(color.FgCyan)
		value = color.New(color.FgYellow)

		side.Print("IP Version: ")
		value.Print(ip.Version, " | ")
		side.Print("IHL: ")
		value.Print(ip.IHL, "    | ")
		side.Print("TOS: ")
		value.Print(ip.TOS, " | ")
		side.Print("	Length: ")
		value.Print(ip.Length, " 	  	 |\n")
		side.Print("     ID: ")
		value.Print(ip.Id)
		value.Print("            | ")
		side.Print("Flags: ")
		value.Print(ip.Flags, "  | ")
		side.Print("OFF: ")
		value.Print(ip.FragOffset)
		value.Print(" 	   	 | ")
		fmt.Println()
		side.Print("TTL: ")
		value.Print(ip.TTL, "       | ")
		side.Print("Pro: ")
		value.Print(ip.Protocol, "  |")
		side.Print("		CHS: ")
		value.Print(ip.Checksum, "   	   	 | \n")
		side.Print("			Src: ")
		value.Print(ip.SrcIP, "	    	   	 | \n")
		side.Print("	        	Dst: ")
		value.Print(ip.DstIP, " 		   	 |  ")

	}
}
