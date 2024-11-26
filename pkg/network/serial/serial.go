package serial

import (
    "sync"
)

var (
	serial      uint8
	serialMutex sync.Mutex
)

func GetNextSerial() uint8 {
	serialMutex.Lock()
	defer serialMutex.Unlock()

	serial = (serial + 1) % 255
	return serial
}