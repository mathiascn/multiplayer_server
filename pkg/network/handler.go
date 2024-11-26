package network

import (
	"fmt"
	"net"

	"github.com/mathiascn/multiplayer_server/pkg/network/handlers"
	"github.com/mathiascn/multiplayer_server/pkg/protocol"
)

func HandlePacket(conn *net.UDPConn, addr *net.UDPAddr, data []byte) {
	fmt.Printf("{\n\treceived from: %v\n", addr)

	packet, err := protocol.DecodePacket(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	switch packet.MessageType {
	case protocol.MessageTypeHandshake:
		handlers.HandleHandshake(packet, conn, addr)
	default:
		fmt.Printf("Unknown message type %d\n", packet.MessageType)
	}

	fmt.Println("\n}")
}
