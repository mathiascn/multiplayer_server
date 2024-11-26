package messages

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/google/uuid"
	"github.com/mathiascn/multiplayer_server/pkg/constants"
	"github.com/mathiascn/multiplayer_server/pkg/version"
)

type HandshakePayload struct {
	Major      uint8
	Minor      uint8
	Patch      uint8
	ServerType string
	PlayerID   string
}

func NewHandshakePayload() HandshakePayload {
	return HandshakePayload{
		Major:      version.ServerVersionMajor,
		Minor:      version.ServerVersionMinor,
		Patch:      version.ServerVersionPatch,
		ServerType: constants.ServerType,
		PlayerID:   uuid.NewString(),
	}
}

func (h HandshakePayload) Serialize() []byte {
	buffer := new(bytes.Buffer)

	buffer.WriteByte(h.Major)
	buffer.WriteByte(h.Minor)
	buffer.WriteByte(h.Patch)

	// write server bytes length and server type bytes
	serverTypeBytes := []byte(h.ServerType)
	if err := binary.Write(buffer, binary.LittleEndian, uint16(len(serverTypeBytes))); err != nil {
		fmt.Println(err)
	}
	buffer.Write(serverTypeBytes)

	playerIdBytes := []byte(h.PlayerID)
	if err := binary.Write(buffer, binary.LittleEndian, uint16(len(playerIdBytes))); err != nil {
		fmt.Println(err)
	}
	buffer.Write(playerIdBytes)

	return buffer.Bytes()
}

func DeserializeHandshakePayload(data []byte) (HandshakePayload, error) {
	if len(data) < 3 {
		return HandshakePayload{}, fmt.Errorf("invalid handshake payload size")
	}

	buffer := bytes.NewReader(data)

	major, err := buffer.ReadByte()
	if err != nil {
		return HandshakePayload{}, fmt.Errorf("failed to read Major: %v", err)
	}

	minor, err := buffer.ReadByte()
	if err != nil {
		return HandshakePayload{}, fmt.Errorf("failed to read Minor: %v", err)
	}

	patch, err := buffer.ReadByte()
	if err != nil {
		return HandshakePayload{}, fmt.Errorf("failed to read Patch: %v", err)
	}

	return HandshakePayload{
		Major: major,
		Minor: minor,
		Patch: patch,
	}, nil
}
