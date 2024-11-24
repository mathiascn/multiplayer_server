package protocol

import (
	"testing"
	"time"
)

func TestEncodeDecodePacket(t *testing.T) {
	tests := []struct {
		name          string
		input         Packet
		rawInput      []byte // For testing DecodePacket directly
		expected      Packet
		expectedError string // Expected error message or part of it
	}{
		{
			name: "valid packet",
			input: Packet{
				MessageType: 1,
				ErrorFlag:   0,
				Serial:      42,
				Timestamp:   time.Now().UnixMilli(),
				Payload:     []byte("hello world"),
			},
			expected: Packet{
				MessageType: 1,
				ErrorFlag:   0,
				Serial:      42,
				Payload:     []byte("hello world"),
			},
			expectedError: "",
		},
		{
			name:          "packet too short",
			rawInput:      []byte{1, 0, 0},
			expectedError: "packet too short",
		},
		{
			name: "invalid payload length",
			rawInput: []byte{
				1, 0, 0, 0, 5, // Message Type + Payload Size (5 bytes)
				0, 42, // ErrorFlag + Serial
				0, 0, 0, 0, 0, 0, 0, 1, // Timestamp
				1, 2, 3, // Payload (only 3 bytes, mismatch)
			},
			expectedError: "invalid payload length",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.rawInput != nil {
				// Test DecodePacket for raw input
				_, err := DecodePacket(tt.rawInput)
				if err == nil {
					t.Fatalf("expected an error, got nil")
				}
				if !containsError(err.Error(), tt.expectedError) {
					t.Errorf("unexpected error message: got %q, want %q", err.Error(), tt.expectedError)
				}
				return
			}

			// Test EncodePacket for valid input
			encoded, err := EncodePacket(tt.input)
			if tt.expectedError != "" {
				if err == nil {
					t.Fatalf("expected an error, got nil")
				}
				if !containsError(err.Error(), tt.expectedError) {
					t.Errorf("unexpected error message: got %q, want %q", err.Error(), tt.expectedError)
				}
				return
			}

			// Test DecodePacket for encoded data
			decoded, err := DecodePacket(encoded)
			if err != nil {
				t.Fatalf("failed to decode packet: %v", err)
			}

			if decoded.MessageType != tt.expected.MessageType {
				t.Errorf("MessageType mismatch: got %d, want %d", decoded.MessageType, tt.expected.MessageType)
			}
			if decoded.ErrorFlag != tt.expected.ErrorFlag {
				t.Errorf("ErrorFlag mismatch: got %d, want %d", decoded.ErrorFlag, tt.expected.ErrorFlag)
			}
			if decoded.Serial != tt.expected.Serial {
				t.Errorf("Serial mismatch: got %d, want %d", decoded.Serial, tt.expected.Serial)
			}
			if string(decoded.Payload) != string(tt.expected.Payload) {
				t.Errorf("Payload mismatch: got %s, want %s", decoded.Payload, tt.expected.Payload)
			}

			if decoded.Timestamp < tt.input.Timestamp || decoded.Timestamp > time.Now().UnixMilli() {
				t.Errorf("Timestamp mismatch: got %d, expected close to %d", decoded.Timestamp, tt.input.Timestamp)
			}
		})
	}
}

func containsError(actual, expected string) bool {
	return actual == expected || len(expected) > 0 && containsSubstring(actual, expected)
}

func containsSubstring(actual, expected string) bool {
	return len(actual) >= len(expected) && actual[:len(expected)] == expected
}
