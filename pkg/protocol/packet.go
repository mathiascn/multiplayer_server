package protocol


import (
	"fmt"
)


func EncodePacket(messageType MessageType, payload string) []byte {
	// First byte is the message type
	header := []byte{byte(messageType)}

	// Next 4 bytes are the payload length in Big-Endian
	// https://en.wikipedia.org/wiki/Endianness
	length := uint32(len(payload))
	header = append(header, byte(length>>24), byte(length>>16), byte(length>>8), byte(length))

	// Combine header and payload into a byte slice
	packet := append(header, []byte(payload)...)
	return packet
}


func DecodePacket(packet []byte) (MessageType, string, error) {
	fmt.Println("reading packet of size:", len(packet))
	// If packet is less than 5 it does not contain a complete header
	if len(packet) < 5 {
		return 0, "", fmt.Errorf("packet too short")
	}

	messageType := MessageType(packet[0])
	length := uint32(packet[1])<<24 | uint32(packet[2])<<16 | uint32(packet[3])<<8 | uint32(packet[4])

	// Making sure the payload length matches the remaining packet size
	if int(length) != len(packet)-5 {
		return 0, "", fmt.Errorf("invalid payload length: expected %d, got %d", length, len(packet)-5)
	}

	// Extract the payload
	payload := string(packet[5:]) // Convert the payload to a string

	return messageType, payload, nil
}