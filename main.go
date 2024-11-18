package main

import (
	"flag"
	"fmt"
	"net"
)

type MessageType byte

const (
	MessageTypeHandshake MessageType = iota // 0
	MessageTypeMove                         // 1
	MessageTypeShoot                        // 2
	MessageTypeChat                         // 3
)

func createPacket(messageType MessageType, payload string) []byte {
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

func readPacket(packet []byte) (MessageType, string, error) {
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

func handlePacket(conn *net.UDPConn,addr *net.UDPAddr, packet []byte) {
	_, message, err := readPacket(packet) // Use message type
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Received from %v: %s\n", addr, message)

	response := fmt.Sprintf("Echo: %s", message)
	newPacket := createPacket(MessageTypeHandshake, response)

	_, err = conn.WriteToUDP(newPacket, addr)
	if err != nil {
		fmt.Println("Error sending response to client:", err)
	}
}


func main() {
	// Flag reads ip and port from CLI
	ip := flag.String("ip", "0.0.0.0", "IP address to bind to")
	port := flag.Int("port", 8080, "Port to listen on")
	flag.Parse()

	parsedIP := net.ParseIP(*ip)
	if parsedIP == nil {
        fmt.Println("Error: Invalid IP address format")
        return
    }

	fmt.Printf("Server will start on IP: %s, Port: %d\n", *ip, *port)
	addr := net.UDPAddr{
		Port: *port,
		IP:   net.ParseIP(*ip),
	}

	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Println("Error starting UDP server:", err)
		return
	}
	defer conn.Close()

	fmt.Printf("UDP server is listening on port %d...", *port)
	buffer := make([]byte, 1024)

	for {
		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading from UDP connection:", err)
			continue
		}
		go handlePacket(conn, clientAddr, buffer[:n])
	}
}


