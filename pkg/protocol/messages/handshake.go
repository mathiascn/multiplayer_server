package protocol

type HandshakePayload struct {
	Major uint8
	Minor uint8
	Patch uint8
}

func (h HandshakePayload) Serialize() []byte {
	return []byte{h.Major, h.Minor, h.Patch}
}

func DeserializeHandshakePayload(data []byte) HandshakePayload {
	return HandshakePayload{
		Major: data[0],
		Minor: data[1],
		Patch: data[2],
	}
}