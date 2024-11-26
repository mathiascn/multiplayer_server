package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/mathiascn/multiplayer_server/pkg/network"
	"github.com/mathiascn/multiplayer_server/pkg/constants"
)

func main() {
	// Flag reads ip and port from CLI
	ip := flag.String("ip", constants.Ip, "IP address to bind to")
	port := flag.Int("port", constants.Port, "Port to listen on")
	tickrate := flag.Int("tickrate", constants.Tickrate, "Tickrate of the server")
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

	// start server with default tickrate of 1 billion nanoseconds / tickrate
	// (default 64 tickrate: 15.625 ms per tick)
	server, err := network.NewUDPServer(&addr, network.HandlePacket, time.Duration(1_000_000_000 / *tickrate))
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	fmt.Printf("Server listening on %s\n", addr.String())
	server.Run()
}
