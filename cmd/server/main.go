package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"github.com/mathiascn/multiplayer_server/pkg/network"
)

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

	addr := net.UDPAddr{
		Port: *port,
		IP:   net.ParseIP(*ip),
	}

	server, err := network.NewUDPServer(&addr, network.HandlePacket)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	fmt.Printf("Server listening on %s\n", addr.String())
	server.Run()
}


