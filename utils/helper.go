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

		side.Print("SRC: ")
		value.Print(ethernetPacket.SrcMAC, "     |")

		// Ethernet type is typically IPv4 but could be ARP or other
		side.Print("  Type: ")
		value.Print(ethernetPacket.EthernetType)

		fmt.Println('\n')
	}
}
