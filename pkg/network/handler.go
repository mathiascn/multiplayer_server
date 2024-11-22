package network

import (
	"fmt"
	"net"

	"github.com/mathiascn/multiplayer_server/pkg/protocol"
	"github.com/mathiascn/multiplayer_server/pkg/version"
	"github.com/mathiascn/multiplayer_server/pkg/protocol/messages"
)

func HandleHandshake(message []byte) {
	//deserialize incoming handshake
	handshake, err := messages.DeserializeHandshakePayload([]byte(message))
	if err != nil {
		fmt.Println("Error deserializing handshake", err)
	}

	fmt.Printf("Client handshake: Version %d.%d.%d\n", handshake.Major, handshake.Minor, handshake.Patch)
	//conditional response
	if !version.IsClientCompatible(handshake.Major, handshake.Minor, handshake.Patch) {
		response = ""
	}

	newPacket := protocol.EncodePacket(protocol.MessageTypeHandshake, response)
	_, err = conn.WriteToUDP(newPacket, addr)
	if err != nil {
		fmt.Println("Error sending response to client:", err)
	}
}

func HandlePacket(conn *net.UDPConn,addr *net.UDPAddr, packet []byte) {
	messageType, message, err := protocol.DecodePacket(packet) // Use message type
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Received from %v: %s\n", addr, message)

	switch messageType {
	case protocol.MessageTypeHandshake:
		HandleHandshake()
	default:
		fmt.Println("Unknown message type %d\n", messageType)
	}
}