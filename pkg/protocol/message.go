package protocol

type MessageType byte

const (
	MessageTypeHandshake MessageType = iota // 0
	MessageTypePing                         // 1
	MessageTypePong                         // 2
	MessageTypeMove                         // 3
	MessageTypeShoot                        // 4
	MessageTypeChat                         // 5
)