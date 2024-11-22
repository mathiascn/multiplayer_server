package messages

import (
	"fmt"
	"github.com/mathiascn/multiplayer_server/pkg/version"
)

type HandshakePayload struct {
	Major uint8
	Minor uint8
	Patch uint8
}

func NewHandshakePayload() HandshakePayload {
	return HandshakePayload{
		Major: version.ServerVersionMajor,
		Minor: version.ServerVersionMinor,
		Patch: version.ServerVersionPatch,
	}
}

func (h HandshakePayload) Serialize() []byte {
	return []byte{h.Major, h.Minor, h.Patch}
}

func DeserializeHandshakePayload(data []byte) (HandshakePayload, error) {
	if len(data) < 3 {
		return HandshakePayload{}, fmt.Errorf("invalid handshake payload size")
	}
	return HandshakePayload{
		Major: data[0],
		Minor: data[1],
		Patch: data[2],
	}, nil
}
