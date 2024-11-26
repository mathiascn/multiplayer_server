package handlers

import (
	"fmt"
	"net"

	"github.com/mathiascn/multiplayer_server/pkg/network/serial"
	"github.com/mathiascn/multiplayer_server/pkg/protocol"
	"github.com/mathiascn/multiplayer_server/pkg/protocol/messages"
	"github.com/mathiascn/multiplayer_server/pkg/version"
)

func HandleHandshake(packet protocol.Packet, conn *net.UDPConn, addr *net.UDPAddr) error {
	response := messages.HandshakePayload{}
	errorFlag := 1

	//deserialize incoming handshake
	handshake, err := messages.DeserializeHandshakePayload(packet.Payload)
	if err != nil {
		fmt.Println("Error deserializing handshake", err)
	}
	fmt.Printf("\thandshake: %d.%d.%d\n", handshake.Major, handshake.Minor, handshake.Patch)

	// return error if client incompatible
	if version.IsClientCompatible(handshake.Major, handshake.Minor, handshake.Patch) {
		response = messages.NewHandshakePayload()
		errorFlag = 0
	}

	// encode response
	serial := serial.GetNextSerial()
	p := protocol.Packet{
		MessageType: protocol.MessageTypeHandshake,
		Payload:     response.Serialize(),
		ErrorFlag:   byte(errorFlag),
		Serial:      byte(serial),
	}
	newPacket, err := protocol.EncodePacket(p)

	if err != nil {
		return fmt.Errorf("failed to encode response: %v", err)
	}

	_, err = conn.WriteToUDP(newPacket, addr)
	if err != nil {
		return fmt.Errorf("Error sending response to client: %v", err)
	}

	return nil
}
