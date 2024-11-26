package network

import (
	"fmt"
	"net"
	"time"
)

type UDPServer struct {
	conn       *net.UDPConn
	handler    func(*net.UDPConn, *net.UDPAddr, []byte)
	tickrate   time.Duration
	packetChan chan packet
}

type packet struct {
	clientAddr *net.UDPAddr
	data       []byte
}

func NewUDPServer(addr *net.UDPAddr, handler func(*net.UDPConn, *net.UDPAddr, []byte), tickrate time.Duration) (*UDPServer, error) {
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return nil, err
	}
	return &UDPServer{
		conn: conn,
		handler: handler,
		tickrate: tickrate,
		packetChan: make(chan packet, 1024),
	}, nil
}

func (s *UDPServer) Run() {
	go func() {
		buffer := make([]byte, 1024)
		for {
			n, clientAddr, err := s.conn.ReadFromUDP(buffer)
			if err != nil {
				fmt.Println("Error reading from UDP:", err)
				continue
			}

			// put packet in the packet channel
			s.packetChan <- packet{clientAddr: clientAddr, data: buffer[:n]}
		}
	}()

	ticker := time.NewTicker(s.tickrate)
	defer ticker.Stop()

	for range ticker.C {
		for len(s.packetChan) > 0 {
			p := <-s.packetChan
			go s.handler(s.conn, p.clientAddr, p.data)
		}

		s.Tick()
	}
}

func (s *UDPServer) Tick() {
	// do game updates here
}

func (s *UDPServer) Close() {
	s.conn.Close()
}
