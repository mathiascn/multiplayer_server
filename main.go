package main

import (
	"flag"
	"fmt"
	"net"
)

func main() {
	//Read ip and port from CLI
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

		message := string(buffer[:n])
		fmt.Printf("Received from %v: %s\n", clientAddr, message)

		response := fmt.Sprintf("Echo: %s", message)
		_, err = conn.WriteToUDP([]byte(response), clientAddr)
		if err != nil {
			fmt.Println("Error sending response to client:", err)
		}
	}
}
