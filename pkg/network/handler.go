package network

import (
	"fmt"
	"net"
	"sync"
	"github.com/mathiascn/multiplayer_server/pkg/protocol"
	"github.com/mathiascn/multiplayer_server/pkg/protocol/messages"
	"github.com/mathiascn/multiplayer_server/pkg/version"
)

var (
	serial uint8
	serialMutex sync.Mutex
)


func getNextSerial() uint8 {
	serialMutex.Lock()
	defer serialMutex.Unlock()

	serial = (serial + 1) % 255
	return serial
}


func HandleHandshake(packet protocol.Packet, conn *net.UDPConn, addr *net.UDPAddr) error {
	response := "Client incompatible"
	errorFlag := 0

	//deserialize incoming handshake
	handshake, err := messages.DeserializeHandshakePayload(packet.Payload)
	if err != nil {
		fmt.Println("Error deserializing handshake", err)
	}
	fmt.Printf("Client handshake: Version %d.%d.%d\n", handshake.Major, handshake.Minor, handshake.Patch)

	// return error if client incompatible
	if version.IsClientCompatible(handshake.Major, handshake.Minor, handshake.Patch) {
		response = "Client compatible"
		errorFlag = 1
	}


	// encode response
	serial = getNextSerial()
	p := protocol.Packet{
		MessageType: protocol.MessageTypeHandshake,
		Payload: []byte(response),
		ErrorFlag: byte(errorFlag),
		Serial: byte(serial),
	}
	newPacket, err := protocol.EncodePacket(p)

	if err != nil {
		return fmt.Errorf("failed to encode response %s", err)
	}

	_, err = conn.WriteToUDP(newPacket, addr)
	if err != nil {
		fmt.Println("Error sending response to client:", err)
	}

	return nil
}


func HandlePacket(conn *net.UDPConn, addr *net.UDPAddr, data []byte) {
	packet, err := protocol.DecodePacket(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Received from %v: %s\n", addr, string(packet.Payload))

	switch packet.MessageType {
	case protocol.MessageTypeHandshake:
		HandleHandshake(packet, conn, addr)
	default:
		fmt.Printf("Unknown message type %d\n", packet.MessageType)
	}
}