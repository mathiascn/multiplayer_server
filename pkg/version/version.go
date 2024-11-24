package version

import "fmt"

//Server version
const (
	ServerVersionMajor uint8 = 0
	ServerVersionMinor uint8 = 1
	ServerVersionPatch uint8 = 0
)

//Client version
//Minimum requirement for client to be compatible
const (
	ClientVersionMajor uint8 = 0
	ClientVersionMinor uint8 = 1
	ClientVersionPatch uint8 = 1
)

func FormatVersion(major, minor, patch uint8) string {
	return fmt.Sprintf("%d.%d.%d", major, minor, patch)
}

func IsClientCompatible(clientMajor, clientMinor, clientPatch uint8) bool {
	//Major version have breaking changes
	//this must match exactly
	if clientMajor != ServerVersionMajor {
		return false
	}

	// Client minor and patch must be at least the minimum requirement
	if clientMinor < ServerVersionMinor {
		return false
	}

	if clientPatch < ServerVersionPatch {
		return false
	}

	return true
}
