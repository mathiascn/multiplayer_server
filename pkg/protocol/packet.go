// Packet structure:
// [Message Type: 1 byte]
// [Payload Size: 4 bytes]
// [Serial: 1 byte]
// [Error Flag: 1 byte]
// [Timestamp: 8 bytes]
// [Payload: variable length]
package protocol

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"time"
)

const (
	messageTypeSize = 1
	payloadSizeSize = 4
	serialSize      = 1
	errorFlagSize   = 1
	timestampSize   = 8
	headerSize      = messageTypeSize + payloadSizeSize + serialSize + errorFlagSize + timestampSize
)

var (
	order = binary.LittleEndian
)

type Packet struct {
	MessageType MessageType
	ErrorFlag   byte
	Serial      uint8
	Timestamp   int64
	Payload     []byte
}

func EncodePacket(packet Packet) ([]byte, error) {
	buffer := new(bytes.Buffer)

	// write message type (1 byte)
	buffer.Write([]byte{byte(packet.MessageType)})

	// write payload size (4 bytes)
	payloadSize := uint32(len(packet.Payload))
	if err := binary.Write(buffer, order, payloadSize); err != nil {
		return nil, err
	}

	// write error flag (1 byte)
	buffer.Write([]byte{packet.ErrorFlag})

	// write serial (1 byte)
	buffer.Write([]byte{packet.Serial})

	// write unix timestamp (8 bytes)
	if packet.Timestamp == 0 {
		packet.Timestamp = time.Now().UnixMilli()
	}
	if err := binary.Write(buffer, order, packet.Timestamp); err != nil {
		return nil, err
	}

	// write payload (variable bytes)
	buffer.Write(packet.Payload)

	return buffer.Bytes(), nil
}

func DecodePacket(data []byte) (Packet, error) {
	dataLength := len(data)
	fmt.Println("\tpacket size:", dataLength)

	// return error if length is less than header size
	if dataLength < headerSize {
		return Packet{}, fmt.Errorf("packet too short")
	}

	buffer := bytes.NewReader(data)

	// read message type (1 byte)
	messageType, err := buffer.ReadByte()
	if err != nil {
		return Packet{}, err
	}

	// read payload size (4 bytes)
	var payloadSize uint32
	if err := binary.Read(buffer, order, &payloadSize); err != nil {
		return Packet{}, err
	}

	// ensuring packet length is of expected size
	expectedLength := headerSize + int(payloadSize)
	if dataLength < expectedLength {
		return Packet{}, fmt.Errorf("invalid payload length: expected %d, got %d", expectedLength, dataLength)
	}

	// read error flag (1 byte)
	errorFlag, err := buffer.ReadByte()
	if err != nil {
		return Packet{}, err
	}

	// read serial (1 byte)
	serial, err := buffer.ReadByte()
	if err != nil {
		return Packet{}, err
	}

	// read timestamp (8 bytes)
	var timestamp int64
	if err := binary.Read(buffer, order, &timestamp); err != nil {
		return Packet{}, err
	}

	// read payload
	payload := make([]byte, payloadSize)
	if _, err := buffer.Read(payload); err != nil {
		return Packet{}, err
	}

	// fmt.Printf("MessageType: %d\n", messageType)
	// fmt.Printf("ErrorFlag: %d\n", errorFlag)
	// fmt.Printf("Serial: %d\n", serial)
	// fmt.Printf("Timestamp: %d\n", timestamp)
	// fmt.Printf("Payload: %s\n", string(payload))

	return Packet{
		MessageType: MessageType(messageType),
		ErrorFlag:   errorFlag,
		Serial:      serial,
		Timestamp:   timestamp,
		Payload:     payload,
	}, nil
}
