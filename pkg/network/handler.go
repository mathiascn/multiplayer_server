package network

import (
	"net"
	"fmt"
	"github.com/mathiascn/multiplayer_server/pkg/protocol"
)

func HandlePacket(conn *net.UDPConn,addr *net.UDPAddr, packet []byte) {
	_, message, err := protocol.DecodePacket(packet) // Use message type
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Received from %v: %s\n", addr, message)

	response := fmt.Sprintf("Pong: %s", message)
	newPacket := protocol.EncodePacket(protocol.MessageTypePong, response)

	_, err = conn.WriteToUDP(newPacket, addr)
	if err != nil {
		fmt.Println("Error sending response to client:", err)
	}
}